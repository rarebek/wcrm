package v1

import (
	_ "api-gateway/api/docs"
	"context"
	"net/http"
	"time"

	"api-gateway/api/models"
	pbu "api-gateway/genproto/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary			 Create Owner
// @Security 		 ApiKeyAuth
// @Description 	 Api for create owner
// @Tags			 User
// @Accept			 json
// @Produce			 json
// @Param 			 Product body models.CreateOwner true "Create Owner"
// @Success			 200 {object} models.Owner
// @Failure			 404 {object} models.StandartError
// @Failure			 500 {object} models.StandartError
// @Router 			 /v1/owner/create [POST]
func (h HandlerV1) CreateOwner(c *gin.Context) {
	var (
		body        models.CreateOwner
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error marshal": err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
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

// @Summary 		Get Owner
// @Security 		ApiKeyAuth
// @Description 	Api for get product
// @Tags 			User
// @Accept			json
// @Produce 		json
// @Param 			id path string true "Id Owner"
// @Success 		200 {object} models.Owner
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/owner/get/{id} [GET]
func (h *HandlerV1) GetOwner(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	filter := map[string]string{
        "id": id,
    }

	response, err := h.Service.UserService().GetOwner(ctx, &pbu.GetOwnerRequest{
		Filter: filter,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary 		Update Owner
// @Security 		ApiKeyAuth
// @Description 	Api for update product
// @Tags User
// @Accept 			json
// @Produce 		json
// @Param 			Owner body models.UpdateOwner true "Update Owner"
// @Success 		200 {object} models.Owner
// @Failure 		400 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/owner/update [PUT]
func (h *HandlerV1) UpdateOwner(c *gin.Context) {
	var (
		body        models.UpdateOwner
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
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

// @Summary			Delete Owner
// @Security 		ApiKeyAuth
// @Description 	Api for delete product
// @Tags 			User
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Id Owner"
// @Success 		200 {object} models.CheckResponse
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/owner/delete/{id} [DELETE]
func (h *HandlerV1) DeleteOwner(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	filter := map[string]string{
        "id": id,
    }

	response, err := h.Service.UserService().DeleteOwner(ctx, &pbu.GetOwnerRequest{
		Filter: filter,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary  		 Get List Owner
// @Security		 ApiKeyAuth
// @Description 	 Api for get all product
// @Tags			 User
// @Accept			 json
// @Produce			 json
// @Param 			 page path string true "Page Owner"
// @Param			 limit path string true "Limit Owner"
// @Success			 200 {object} models.OwnerList
// @Failure			 404 {object} models.StandartError
// @Failure			 500 {object} models.StandartError
// @Router			 /v1/owners/get/{page}/{limit} [GET]
func (h *HandlerV1) ListOwner(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	page := c.Param("page")
	limit := c.Param("limit")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.UserService().ListOwner(ctx, &pbu.GetAllOwnerRequest{
		Page:  cast.ToInt64(page),
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
