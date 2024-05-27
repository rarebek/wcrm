package v1

import (
	"bytes"
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
	"text/template"

	redis "api-gateway/internal/pkg/redis"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/k0kubun/pp"
	"google.golang.org/protobuf/encoding/protojson"
)

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

	type PageData struct {
		OTP string
	}
	tpl := template.Must(template.ParseFiles("index.html"))
	data := PageData{
		OTP: code,
	}
	var buf bytes.Buffer
	tpl.Execute(&buf, data)
	htmlContent := buf.Bytes()

	auth := smtp.PlainAuth("", "nodirbekgolang@gmail.com", "jgbu bsru qcha buko", "smtp.gmail.com")
	err = smtp.SendMail("smtp.gmail.com:587", auth, "nodirbekgolang@gmail.com", []string{body.Email}, []byte("To: "+body.Email+"\r\nSubject: Email verification\r\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"+string(htmlContent)))

	if err != nil {
		pp.Println(err)
		h.Logger.Error("error writing redis")
		return
	}

	err = redisClient.Client.Set(context.Background(), code, byteData, time.Minute*3).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:  ErrorCodeInternalServerError,
			Error: err.Error(),
		})
		h.Logger.Error("cannot set redis")
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
		Id:          createdOwner.Id,
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

// @Summary 		Login owner
// @Description 	Api for Login
// @Tags 			Register
// @Accept 			json
// @Produce 		json
// @Param 			email query string true "EMAIL"
// @Param 			password query string true "PASSWORD"
// @Success 		200 {object} models.ResponseAccessToken
// @Failure 		400 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/login [get]
func (h *HandlerV1) LogIn(c *gin.Context) {

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	email := c.Query("email")
	password := c.Query("password")

	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	filter := map[string]string{
		"email": email,
	}

	response, err := h.Service.UserService().GetOwner(ctx, &pbu.GetOwnerRequest{
		Filter: filter,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect email. Please try again",
		})
		h.Logger.Error(err.Error())
		return
	}

	if !etc.CheckPasswordHash(password, response.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect password. Please try again",
		})
		h.Logger.Error("Incorrect password. Please try again")
		return
	}

	// Create access and refresh tokens JWT
	h.jwthandler = token.JWTHandler{
		Sub:       response.Id,
		Iss:       time.Now().String(),
		Exp:       time.Now().Add(time.Hour * 6).String(),
		Role:      "owner",
		SigninKey: h.Config.SigningKey,
		Timeot:    h.Config.AccessTokenTimout,
	}

	// aksestoken bn refreshtokeni generatsa qiliah
	access, _, err := h.jwthandler.GenerateAuthJWT()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error generating token",
		})
		h.Logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, models.ResponseOwnerLogin{
		AccessToken: access,
		OwnerId:     response.Id,
	})
}

// @Summary     	Login worker
// @Description     Api for Login
// @Tags       		Register
// @Accept       	json
// @Produce     	json
// @Param       	Owner body  models.LoginWorker true "worker"
// @Success     	200 {object} models.ResponseAccessToken
// @Failure     	400 {object} models.StandartError
// @Failure     	500 {object} models.StandartError
// @Router       	/v1/worker/login [POST]
func (h *HandlerV1) LogInWorker(c *gin.Context) {

	var (
		body        models.LoginWorker
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	filter := map[string]string{
		"login_key": body.LoginKey,
	}

	resWorker, err := h.Service.UserService().GetWorker(ctx, &pbu.GetWorkerRequest{
		Filter: filter,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect login key. Please try again",
		})
		h.Logger.Error(err.Error())
		return
	}

	filter = map[string]string{
		"id": resWorker.OwnerId,
	}

	resOwner, err := h.Service.UserService().GetOwner(ctx, &pbu.GetOwnerRequest{
		Filter: filter,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Geting owner error",
		})
		h.Logger.Error(err.Error())
		return
	}

	if resOwner.CompanyName != body.CompanyName {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect company name. Please try again",
		})
		h.Logger.Error("Incorrect company name. Please try again")
		return
	}

	if body.Password != resWorker.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect password. Please try again",
		})
		h.Logger.Error("Incorrect password. Please try again")
		return
	}

	// Create access and refresh tokens JWT
	h.jwthandler = token.JWTHandler{
		Sub:       resWorker.Id,
		Iss:       time.Now().String(),
		Exp:       time.Now().Add(time.Hour * 6).String(),
		Role:      "worker",
		SigninKey: h.Config.SigningKey,
		Timeot:    h.Config.AccessTokenTimout,
	}

	// aksestoken bn refreshtokeni generatsa qiliah
	access, _, err := h.jwthandler.GenerateAuthJWT()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error generating token",
		})
		h.Logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, models.ResponseAccessToken{
		AccessToken: access,
		WorkerId:    resWorker.Id,
	})
}
