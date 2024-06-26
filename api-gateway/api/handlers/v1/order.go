package v1

import (
	_ "api-gateway/api/docs"
	"api-gateway/api/models"
	pbo "api-gateway/genproto/order"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary Create Order
// @Security ApiKeyAuth
// @Description Api for create order
// @Tags Order
// @Accept json
// @Produce json
// @Param Order body models.CreateOrder true "Create Order"
// @Success 200 {object} models.Order
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/order/create [POST]
func (h *HandlerV1) CreateOrder(c *gin.Context) {
	var (
		body        models.CreateOrder
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

	protoProducts := make([]*pbo.ProductCheck, len(body.Products))
	for i, product := range body.Products {
		protoProducts[i] = &pbo.ProductCheck{
			Title: product.Title,
			Count: product.Count,
			Price: product.Price,
		}
	}

	response, err := h.Service.OrderService().CreateOrder(ctx, &pbo.Order{
		TableNumber: body.TableNumber,
		WorkerId:    body.WorkerId,
		Products:    protoProducts,
		Tax:         body.Tax,
		TotalPrice:  body.TotalPrice,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	products := make([]models.Product, len(body.Products))
	for i, p := range body.Products {
		products[i] = models.Product{
			Title: p.Title,
			Price: p.Price,
		}
	}

	// Store products in memory for simplicity. You can use a database for persistence.
	h.ProductStore = append(h.ProductStore, products...)

	c.JSON(http.StatusCreated, response)
}

// @Summary Get Products
// @Description Get the list of products
// @Tags Product
// @Produce json
// @Success 200 {array} models.Product
// @Router /v1/products/bot [GET]
func (h *HandlerV1) GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, h.ProductStore)

	h.ProductStore = nil
}

// @Summary Delete Products
// @Description deleted bot product
// @Tags Product
// @Produce json
// @Success 200 {array} models.Product
// @Router /v1/products/bot [DELETE]
func (h *HandlerV1) DeleteProductBot(c *gin.Context) {
	h.ProductStore = nil
}

// @Summary 		Get Order
// @Security 		ApiKeyAuth
// @Description 	Api for get order
// @Tags 			Order
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Id Order"
// @Success 		200 {object} models.GetOrderResponse
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/order/get/{id} [GET]
func (h *HandlerV1) GetOrder(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.OrderService().GetOrder(ctx, &pbo.OrderId{
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

// @Summary 		Update Order
// @Security 		ApiKeyAuth
// @Description 	Api for update order
// @Tags 			Order
// @Accept 			json
// @Produce 		json
// @Param 			Order body models.UpdateOrder true "Update Order"
// @Succes 			200 {object} models.Order
// @Failure 		400 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/order/update [PUT]
func (h *HandlerV1) UpdateOrder(c *gin.Context) {
	var (
		body        models.UpdateOrder
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

	response, err := h.Service.OrderService().UpdateOrder(ctx, &pbo.Order{
		Id:         body.Id,
		Tax:        body.Tax,
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

// @Summary 		Delete Order
// @Security 		ApiKeyAuth
// @Description 	Api for delete order
// @Tags 			Order
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Id Order"
// @Success 		200 {object} models.CheckResponse
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/order/delete/{id} [DELETE]
func (h *HandlerV1) DeleteOrder(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.OrderService().DeleteOrder(ctx, &pbo.OrderId{
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

// @Summary 		Get List Order
// @Security 		ApiKeyAuth
// @Description 	Api for get all product
// @Tags			Order
// @Accept 			json
// @Produce 		json
// @Param 			page path string true "Page Order"
// @Param 			limit path string true "Limit Order"
// @Param 			worker-id path string true "Worker Id"
// @Success 		200 {object} models.OrderList
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/orders/get/{page}/{limit}/{worker-id} [GET]
func (h *HandlerV1) ListOrder(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	page := c.Param("page")
	limit := c.Param("limit")
	workerId := c.Param("worker-id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.OrderService().GetOrders(ctx, &pbo.GetAllOrderRequest{
		Page:     cast.ToInt64(page),
		Limit:    cast.ToInt64(limit),
		WorkerId: workerId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println(response)

	c.JSON(http.StatusOK, response)
}
