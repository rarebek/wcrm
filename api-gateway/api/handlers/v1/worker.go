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
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
)

// Create Worker
// @Summary Create Worker
// @Description Api for create worker
// @Tags Worker
// @Accept json
// @Produce json
// @Param Worker body models.CreateWorker true "Create Worker"
// @Success 200 {object} models.Worker
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/worker/create [POST]
func (h HandlerV1) CreateWorker(c *gin.Context) {
	var (
		body        models.CreateWorker
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	id := uuid.New().String()

	response, err := h.Service.UserService().CreateWorker(ctx, &pbu.Worker{
		Id:       id,
		FullName: body.FullName,
		LoginKey: body.LoginKey,
		Password: body.Password,
		OwnerId:  body.OwnerId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// Get Worker
// @Summary Get Worker
// @Description Api for get worker
// @Tags Worker
// @Accept json
// @Produce json
// @Param id path string true "Id Worker"
// @Success 200 {object} models.Worker
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/worker/get/{id} [GET]
func (h *HandlerV1) GetWorker(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := h.Service.UserService().GetWorker(ctx, &pbu.GetWorkerRequest{
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

// Update Worker
// @Summary Update Worker
// @Description Api for update worker
// @Tags Worker
// @Accept json
// @Produce json
// @Param Worker body models.UpdateProduct true "Update Worker"
// @Success 200 {object} models.Worker
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/worker/update [PUT]
func (h *HandlerV1) UpdateWorker(c *gin.Context) {
	var (
		body        models.UpdateWorker
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

	response, err := h.Service.UserService().UpdateWorker(ctx, &pbu.Worker{
		Id:        body.Id,
		FullName:  body.FullName,
		LoginKey:  body.LoginKey,
		Password:  body.Password,
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

// Delete Worker
// @Summary Delete Worker
// @Description Api for delete worker
// @Tags Worker
// @Accept json
// @Produce json
// @Param id path string true "Id Worker"
// @Success 200 {object} models.CheckResponse
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/worker/delete/{id} [DELETE]
func (h *HandlerV1) DeleteWorker(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := h.Service.UserService().DeleteWorker(ctx, &pbu.GetWorkerRequest{
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

// Get List Worker
// @Summary Get List Worker
// @Description Api for get all worker
// @Tags Worker
// @Accept json
// @Produce json
// @Param page path string true "Page Worker"
// @Param limit path string true "Limit Worker"
// @Success 200 {object} models.WorkerList
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/workers/get/{page}/{limit} [GET]
func (h *HandlerV1) ListWorker(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	page := c.Param("page")
	limit := c.Param("limit")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := h.Service.UserService().ListWorker(ctx, &pbu.GetAllWorkerRequest{
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
