package v1

import (
	"context"
	"encoding/json"

	"net/http"
	"net/smtp"

	"strings"
	"time"

	"api-gateway/api/models"
	pbu "api-gateway/genproto/user"
	etc "api-gateway/internal/pkg/etc"
	l "api-gateway/internal/pkg/logger"
	token "api-gateway/internal/pkg/token"

	redis "api-gateway/internal/pkg/redis"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
)

// jwt

// @Summary 		Register
// @Description 	Api for Registering
// @Tags 			Register
// @Accept 			json
// @Produce 		json
// @Param 			Owner body  models.RegisterOwner true "owner"
// @Success 		200 {object} string
// @Failure 		400 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/register [post]
func (h *HandlerV1) Register(c *gin.Context) {
	var (
		body        models.RegisterOwner
		jsonMarshal protojson.MarshalOptions
	)

	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("failed to bind json", l.Error(err))
		return
	}

	body.Email = strings.TrimSpace(body.Email)
	body.Email = strings.ToLower(body.Email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	if err = body.IsEmail(); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Incorrect email format",
		})
		h.Logger.Error("Incorrect email format", l.Error(err))
		return
	}

	if err = body.IsComplexPassword(); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Password must be at least 8 characters long and contain both upper and lower case letters",
		})
		h.Logger.Error("password format incorrect", l.Error(err))
		return
	}

	exists, err := h.Service.UserService().CheckFieldOwner(ctx, &pbu.CheckFieldRequest{
		Field: "email",
		Value: body.Email,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:  ErrorCodeInternalServerError,
			Error: err.Error(),
		})
		h.Logger.Error("failed to check uniqueness: ", l.Error(err))
		return
	}

	if exists.Exist {
		c.JSON(http.StatusConflict, models.ResponseError{
			Code:  ErrorCodeAlreadyExists,
			Error: "this email is already in use",
		})
		h.Logger.Error("email is already exist in database")
		return
	}

	byteData, err := json.Marshal(body)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Code:  ErrorCodeInternalServerError,
			Error: err.Error(),
		})
		h.Logger.Error("failed while marshalling user data")
		return
	}
	redisClient, err := redis.New(h.Config)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:  ErrorCodeInternalServerError,
			Error: "error opening redis",
		})
		h.Logger.Error("error opening redis")
		return
	}

	code := etc.StringNumber(6)

	err = redisClient.Client.Set(context.Background(), code, byteData, time.Minute*3).Err()

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:  ErrorCodeInternalServerError,
			Error: err.Error(),
		})
		h.Logger.Error("cannot set redis")
		return
	}

	auth := smtp.PlainAuth("", "asadfaxriddinov611@gmail.com", "drkeagdlwrfanrdp", "smtp.gmail.com")
	err = smtp.SendMail("smtp.gmail.com:587", auth, "asadfaxriddinov611@gmail.com", []string{body.Email}, []byte(code))

	if err != nil {
		h.Logger.Error("error writing redis")
		return
	}

	c.JSON(http.StatusOK, models.RegisterOwnerResponse{
		Message: "One time verification password sent to your email. Please verify",
	})
}

// @Summary 		Verify owner
// @Description 	LogIn - Verify a user with code sent to their email
// @Tags 			Register
// @Accept 			json
// @Produce 		json
// @Param 			email query string true "Email"
// @Param 			code query string true "Code"
// @Success 		200 {object} models.OwnerResponse
// @Failure 		400 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/verification [get]
func (h *HandlerV1) Verify(c *gin.Context) {

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	email := c.Query("email")
	code := c.Query("code")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	redisClient, err := redis.New(h.Config)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:  ErrorCodeInternalServerError,
			Error: "error opening redis",
		})
		h.Logger.Error("error opening redis")
		return
	}

	val, err := redisClient.Client.Get(ctx, code).Result()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Verification code is expired, try again.",
		})
		h.Logger.Error("Verification code is expired", l.Error(err))
		return
	}

	var owner models.Owner
	if err := json.Unmarshal([]byte(val), &owner); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unmarshiling error",
		})
		h.Logger.Error("error unmarshalling owner", l.Error(err))
		return
	}

	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	if email != owner.Email {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect email. Try again",
		})
		return
	}

	id := uuid.New().String()

	hashPassword, err := etc.HashPassword(owner.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error message",
		})
		h.Logger.Error("Error hashing possword", l.Error(err))
		return
	}

	// Create access and refresh tokens JWT
	h.jwthandler = token.JWTHandler{
		Sub:       id,
		Iss:       time.Now().String(),
		Exp:       time.Now().Add(time.Hour * 6).String(),
		Role:      "owner",
		SigninKey: h.Config.SigningKey,
		Timeot:    h.Config.AccessTokenTimout,
	}
	// aksestoken bn refreshtokeni generatsa qiliah
	access, _, err := h.jwthandler.GenerateAuthJWT()

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:  ErrorCodeInternalServerError,
			Error: err.Error(),
		})
		h.Logger.Error("cannot create access and refresh token", l.Error(err))
		return
	}

	createdOwner, err := h.Service.UserService().CreateOwner(ctx, &pbu.Owner{
		Id:          id,
		FullName:    owner.FullName,
		CompanyName: owner.CompanyName,
		Email:       owner.Email,
		Password:    hashPassword,
		Avatar:      owner.Avatar,
		Tax:         owner.Tax,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("cannot create owner", l.Error(err))
		return
	}

	response := &models.OwnerResponse{
		Id:          id,
		FullName:    createdOwner.FullName,
		CompanyName: createdOwner.CompanyName,
		Email:       createdOwner.Email,
		Password:    createdOwner.Password,
		Avatar:      createdOwner.Avatar,
		Tax:         createdOwner.Tax,
		AccessToken: access,
	}

	c.JSON(http.StatusOK, response)
}
