package v1

import (
	"context"
	_ "evrone_service/api_gateway/api/docs"
	"net/http"
	"time"

	// "github.com/casbin/casbin/v2"
	// "evrone_service/api_gateway/api/middleware"
	"evrone_service/api_gateway/api/models"
	pbp "evrone_service/api_gateway/genproto/product"

	// grpcClient "evrone_service/api_gateway/internal/infrastructure/grpc_service_client"
	// "evrone_service/api_gateway/internal/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
)

// Create Product
// @Summary Create Product
// @Description Api for create product
// @Tags Product
// @Accept json
// @Produce json
// @Param Product body models.CreateProduct true "Create Product"
// @Success 200 {object} models.Product
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/product/create [POST]
func (h HandlerV1) CreateProduct(c *gin.Context) {
	var (
		body        models.CreateProduct
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

	response, err := h.Service.ProductService().CreateProduct(ctx, &pbp.Product{
		Title:       body.Title,
		Description: body.Description,
		Price:       body.Price,
		Discount:    body.Discount,
		Picture:     body.Picture,
		CategoryId:  body.CategoryId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusCreated, response)
}


// Get Product
// @Summary Get Product
// @Description Api for get product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Id Product"
// @Success 200 {object} models.Product
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/product/get/{id} [GET]
func (h *HandlerV1) GetProduct(c *gin.Context)  {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := h.Service.ProductService().GetProduct(ctx, &pbp.GetProductRequest{
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

// Update Product 
// @Summary Update Product
// @Description Api for update product
// @Tags Product
// @Accept json
// @Produce json
// @Param Product body models.UpdateProduct true "Update Product"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/product/update [PUT]
func (h *HandlerV1) UpdateProduct(c *gin.Context)  {
	var (
		body      models.UpdateProduct
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

	response, err := h.Service.ProductService().UpdateProduct(ctx, &pbp.Product{
		Id:          body.Id,
		Title:       body.Title,
		Description: body.Description,
		Price:       body.Price,
		Discount:    body.Discount,
		Picture:     body.Picture,
		CategoryId:  body.CategoryId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Delete Product
// @Summary Delete Product
// @Description Api for delete product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Id Product"
// @Success 200 {object} models.CheckResponse
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/product/delete/{id} [DELETE]
func (h *HandlerV1) DeleteProduct(c *gin.Context)  {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := h.Service.ProductService().DeleteProduct(ctx, &pbp.DeleteProductRequest{
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

// Get List Product
// @Summary Get List Product
// @Description Api for get all product
// @Tags Product
// @Accept json
// @Produce json
// @Param page path string true "Page Product"
// @Param limit path string true "Limit Product"
// @Success 200 {object} models.ProductList
// @Failure 404 {object} models.StandartError
// @Failure 500 {object} models.StandartError
// @Router /v1/products/get/{page}/{limit} [GET]
func (h *HandlerV1) ListProduct(c *gin.Context)  {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	page := c.Param("page")
	limit := c.Param("limit")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := h.Service.ProductService().ListProduct(ctx, &pbp.GetAllRequest{
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