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
	"github.com/k0kubun/pp"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary 		Create Worker
// @Security 		ApiKeyAuth
// @Description		Api for create worker
// @Tags			Worker
// @Accept 			json
// @Produce 		json
// @Param 			Worker body models.CreateWorker true "Create Worker"
// @Success 		200 {object} models.Worker
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/worker/create [POST]
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
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

// @Summary 		Get Worker
// @Security 		ApiKeyAuth
// @Description 	Api for get worker
// @Tags 			Worker
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Id Worker"
// @Success 		200 {object} models.Worker
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/worker/get/{id} [GET]
func (h *HandlerV1) GetWorker(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	filter := make(map[string]string)
	filter["id"] = id

	response, err := h.Service.UserService().GetWorker(ctx, &pbu.GetWorkerRequest{
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

// @Summary 		Update Worker
// @Security 		ApiKeyAuth
// @Description 	Api for update worker
// @Tags 			Worker
// @Accept 			json
// @Produce 		json
// @Param 			Worker body models.UpdateWorker true "Update Worker"
// @Success 		200 {object} models.Worker
// @Failure 		400 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/worker/update [PUT]
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.UserService().UpdateWorker(ctx, &pbu.Worker{
		Id:       body.Id,
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

	c.JSON(http.StatusOK, response)
}

// @Summary 		Delete Worker
// @Security 		ApiKeyAuth
// @Description 	Api for delete worker
// @Tags 			Worker
// @Accept			json
// @Produce 		json
// @Param 			id path string true "Id Worker"
// @Success 		200 {object} models.CheckResponse
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/worker/delete/{id} [DELETE]
func (h *HandlerV1) DeleteWorker(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.UserService().DeleteWorker(ctx, &pbu.IdRequest{
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

// @Summary 		Get List Worker
// @Security 		ApiKeyAuth
// @Description 	Api for get all worker
// @Tags 			Worker
// @Accept 			json
// @Produce 		json
// @Param 			page path string true "Page Worker"
// @Param 			limit path string true "Limit Worker"
// @Param 			owner-id path string true "Owner ID od Worker"
// @Success 		200 {object} models.WorkerList
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/workers/get/{page}/{limit}/{owner-id} [GET]
func (h *HandlerV1) ListWorker(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	page := c.Param("page")
	limit := c.Param("limit")
	ownerId := c.Param("owner-id")
	pp.Println("OWNER ID IN API GATEWAY: ", ownerId)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	filter := map[string]string{
		"owner_id": ownerId,
	}

	response, err := h.Service.UserService().ListWorker(ctx, &pbu.GetAllWorkerRequest{
		Page:   cast.ToInt64(page),
		Limit:  cast.ToInt64(limit),
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
