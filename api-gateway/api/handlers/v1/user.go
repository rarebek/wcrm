package v1

import (
	"context"
	_ "evrone_service/api_gateway/api/docs"
	"net/http"
	"time"

	// "github.com/casbin/casbin/v2"
	// "evrone_service/api_gateway/api/middleware"
	"evrone_service/api_gateway/api/models"
	pbu "evrone_service/api_gateway/genproto/user_service"

	// grpcClient "evrone_service/api_gateway/internal/infrastructure/grpc_service_client"
	// "evrone_service/api_gateway/internal/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
)

// Create Owner
// @Summary Create Owner
// @Description Api for create owner
// @Tags Owner
// @Accept json
// @Produce json
// @Param Product body models.CreateOwner true "Create Owner"
// @Success 200 {object} models.Owner
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/owner/create [POST]
func (h HandlerV1) CreateOwner(c *gin.Context) {
	var (
		body        models.CreateOwner
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	id := uuid.New().String()

	response, err := h.Service.UserService().CreateOwner(ctx, &pbu.Owner{
		Id:          id,
		FullName:    body.FullName,
		CompanyName: body.CompanyName,
		Email:       body.Email,
		Password:    body.Password,
		Avatar:      body.Avatar,
		Tax:         body.Tax,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusCreated, response)
}


// Get Owner
// @Summary Get Owner
// @Description Api for get product
// @Tags Owner
// @Accept json
// @Produce json
// @Param id path string true "Id Owner"
// @Success 200 {object} models.Owner
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/owner/get/{id} [GET]
func (h *HandlerV1) GetOwner(c *gin.Context)  {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := h.Service.UserService().GetOwner(ctx, &pbu.GetOwnerRequest{
		Id: id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Update Owner 
// @Summary Update Owner
// @Description Api for update product
// @Tags Owner
// @Accept json
// @Produce json
// @Param Owner body models.UpdateOwner true "Update Owner"
// @Success 200 {object} models.Owner
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/owner/update [PUT]
func (h *HandlerV1) UpdateOwner(c *gin.Context)  {
	var (
		body      models.UpdateOwner
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := h.Service.UserService().UpdateOwner(ctx, &pbu.Owner{
		Id:          body.Id,
		FullName:    body.FullName,
		CompanyName: body.CompanyName,
		Email:       body.Email,
		Password:    body.Password,
		Avatar:      body.Avatar,
		Tax:         body.Tax,
	})
		

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Delete Owner
// @Summary Delete Owner
// @Description Api for delete product
// @Tags Owner
// @Accept json
// @Produce json
// @Param id path string true "Id Owner"
// @Success 200 {object} models.CheckResponse
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/owner/delete/{id} [DELETE]
func (h *HandlerV1) DeleteOwner(c *gin.Context)  {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := h.Service.UserService().DeleteOwner(ctx, &pbu.GetOwnerRequest{
		Id: id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Get List Owner
// @Summary Get List Owner
// @Description Api for get all product
// @Tags Owner
// @Accept json
// @Produce json
// @Param page path string true "Page Owner"
// @Param limit path string true "Limit Owner"
// @Success 200 {object} models.ProductList
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/owners/get [GET]
func (h *HandlerV1) ListOwner(c *gin.Context)  {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	page := c.Param("page")
	limit := c.Param("limit")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := h.Service.UserService().ListOwner(ctx, &pbu.GetAllOwnerRequest{
		Page: cast.ToInt64(page),
		Limit: cast.ToInt64(limit),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}