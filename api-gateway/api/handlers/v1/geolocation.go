package v1

import (
	_ "api-gateway/api/docs"
	"context"
	"net/http"
	"time"

	"api-gateway/api/models"
	pbu "api-gateway/genproto/user"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary 		Create Geolocation
// @Security 		ApiKeyAuth
// @Description 	Api for create geolocation
// @Tags 			Geolocation
// @Accept 			json
// @Produce 		json
// @Param 			Geolocation body models.CreateGeolocation true "Create Geolocation"
// @Success			200 {object} models.Geolocation
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/geolocation/create [POST]
func (h HandlerV1) CreateGeolocation(c *gin.Context) {
	var (
		body        models.CreateGeolocation
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.UserService().CreateGeolocation(ctx, &pbu.Geolocation{
		Latitude:  float32(body.Latitude),
		Longitude: float32(body.Longitude),
		OwnerId:   body.OwnerId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// @Summary 		Get Geolocation
// @Security 		ApiKeyAuth
// @Description 	Api for get geolocation
// @Tags 			Geolocation
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Id Geolocation"
// @Success 		200 {object} models.Geolocation
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/geolocation/get/{id} [GET]
func (h *HandlerV1) GetGeolocation(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.UserService().GetGeolocation(ctx, &pbu.GetGeolocationRequest{
		Id: cast.ToInt64(id),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary 		Update Geolocation
// @Security 		ApiKeyAuth
// @Description		Api for update geolocation
// @Tags 			Geolocation
// @Accept 			json
// @Produce 		json
// @Param 			Geolocation body models.UpdateGeolocation true "Update Geolocation"
// @Success 		200 {object} models.Geolocation
// @Failure 		400 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/geolocation/update [PUT]
func (h *HandlerV1) UpdateGeolocation(c *gin.Context) {
	var (
		body        models.UpdateGeolocation
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

	response, err := h.Service.UserService().UpdateGeolocation(ctx, &pbu.Geolocation{
		Id:        body.Id,
		Latitude:  float32(body.Latitude),
		Longitude: float32(body.Longitude),
		OwnerId:   body.OwnerId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary 		Delete Geolocation
// @Security 		ApiKeyAuth
// @Description 	Api for delete geolocation
// @Tags 			Geolocation
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Id Geolocation"
// @Success 		200 {object} models.CheckResponse
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/geolocation/delete/{id} [DELETE]
func (h *HandlerV1) DeleteGeolocation(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.UserService().GetGeolocation(ctx, &pbu.GetGeolocationRequest{
		Id: cast.ToInt64(id),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary 		Get List Geolocation
// @Security 		ApiKeyAuth
// @Description 	Api for get all geolocation
// @Tags 			Geolocation
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Owner Id"
// @Success 		200 {object} models.GeolocationList
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/geolocations/get/{page}/{limit} [GET]
func (h *HandlerV1) ListGeolocation(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.UserService().ListGeolocation(ctx, &pbu.GetAllGeolocationRequest{
		OwnerId: id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
