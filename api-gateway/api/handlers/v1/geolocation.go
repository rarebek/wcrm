package v1

import (
	"context"
	_ "evrone_service/api_gateway/api/docs"
	"net/http"
	"time"

	// "github.com/casbin/casbin/v2"
	// "evrone_service/api_gateway/api/middleware"
	"evrone_service/api_gateway/api/models"
	pbu "evrone_service/api_gateway/genproto/user"

	// grpcClient "evrone_service/api_gateway/internal/infrastructure/grpc_service_client"
	// "evrone_service/api_gateway/internal/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
)

// Create Geolocation
// @Summary Create Geolocation
// @Description Api for create geolocation
// @Tags Geolocation
// @Accept json
// @Produce json
// @Param Geolocation body models.CreateProduct true "Create Geolocation"
// @Success 200 {object} models.Geolocation
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/geolocation/create [POST]
func (h HandlerV1) CreateGeolocation(c *gin.Context) {
	var (
		body        models.CreateGeolocation
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


// Get Geolocation
// @Summary Get Geolocation
// @Description Api for get geolocation
// @Tags Geolocation
// @Accept json
// @Produce json
// @Param id path string true "Id Geolocation"
// @Success 200 {object} models.Geolocation
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/geolocation/get/{id} [GET]
func (h *HandlerV1) GetGeolocation(c *gin.Context)  {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
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

// Update Geolocation 
// @Summary Update Geolocation
// @Description Api for update geolocation
// @Tags Geolocation
// @Accept json
// @Produce json
// @Param Geolocation body models.UpdateProduct true "Update Geolocation"
// @Success 200 {object} models.Geolocation
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/geolocation/update [PUT]
func (h *HandlerV1) UpdateGeolocation(c *gin.Context)  {
	var (
		body      models.UpdateGeolocation
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

// Delete Geolocation
// @Summary Delete Geolocation
// @Description Api for delete geolocation
// @Tags Geolocation
// @Accept json
// @Produce json
// @Param id path string true "Id Geolocation"
// @Success 200 {object} models.CheckResponse
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/geolocation/delete/{id} [DELETE]
func (h *HandlerV1) DeleteGeolocation(c *gin.Context)  {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
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

// Get List Geolocation
// @Summary Get List Geolocation
// @Description Api for get all geolocation
// @Tags Geolocation
// @Accept json
// @Produce json
// @Param id path string true "Owner Id"
// @Success 200 {object} models.GeolocationList
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/geolocations/get{page}/{limit} [GET]
func (h *HandlerV1) ListGeolocation(c *gin.Context)  {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
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