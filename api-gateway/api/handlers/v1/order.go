package v1

import (
	"context"
	_ "evrone_service/api_gateway/api/docs"
	"net/http"
	"time"

	// "github.com/casbin/casbin/v2"
	// "evrone_service/api_gateway/api/middleware"
	"evrone_service/api_gateway/api/models"
	pbo "evrone_service/api_gateway/genproto/order"

	// grpcClient "evrone_service/api_gateway/internal/infrastructure/grpc_service_client"
	// "evrone_service/api_gateway/internal/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	// "github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
)

// Create Order
// @Summary Create Order
// @Description Api for create oder
// @Tags Order
// @Accept json
// @Produce json
// @Param Order body models.CreateOrder true "Create Order"
// @Success 200 {object} models.Order
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/order/create [POST]
func (h HandlerV1) CreateOrder(c *gin.Context) {
	var (
		body        models.CreateOrder
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

	response, err := h.Service.OrderService().CreateOrder(ctx, &pbo.Order{
		WorkerId:   body.WorkerId,
		ProductId:  body.ProductId,
		Tax:        body.Tax,
		Discount:   body.Discount,
		TotalPrice: body.TotalPrice,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusCreated, response)
}


// Get Order
// @Summary Get Order
// @Description Api for get order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "Id Order"
// @Success 200 {object} models.Order
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/order/get/{id} [GET]
func (h *HandlerV1) GetOrder(c *gin.Context)  {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := h.Service.OrderService().GetOrder(ctx, &pbo.OrderId{
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

// Update Order 
// @Summary Update Order
// @Description Api for update order
// @Tags Order
// @Accept json
// @Produce json
// @Param Order body models.UpdateOrder true "Update Order"
// @Success 200 {object} models.Order
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/order/update [PUT]
func (h *HandlerV1) UpdateOrder(c *gin.Context)  {
	var (
		body      models.UpdateOrder
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

	response, err := h.Service.OrderService().UpdateOrder(ctx, &pbo.Order{
		Id:         body.Id,
		Tax:        body.Tax,
		Discount:   body.Discount,
		TotalPrice: body.TotalPrice,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Delete Order
// @Summary Delete Order
// @Description Api for delete order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "Id Order"
// @Success 200 {object} models.CheckResponse
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/order/delete/{id} [DELETE]
func (h *HandlerV1) DeleteOrder(c *gin.Context)  {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := h.Service.OrderService().DeleteOrder(ctx, &pbo.OrderId{
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

// Get List Order
// @Summary Get List Order
// @Description Api for get all product
// @Tags Order
// @Accept json
// @Produce json
// @Param page path string true "Page Order"
// @Param limit path string true "Limit Order"
// @Success 200 {object} models.OrderList
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/orders/get [GET]
func (h *HandlerV1) ListOrder(c *gin.Context)  {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	page := c.Param("page")
	limit := c.Param("limit")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := h.Service.OrderService().GetOrders(ctx, &pbo.GetAllOrderRequest{
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