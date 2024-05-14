package v1

import (
	_ "api-gateway/api/docs"
	"context"
	"net/http"
	"time"

	"api-gateway/api/models"
	pbp "api-gateway/genproto/product"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary 		Create Product
// @Security 		ApiKeyAuth
// @Description 	Api for create product
// @Tags 			Product
// @Accept 			json
// @Produce 		json
// @Param 			Product body models.CreateProduct true "Create Product"
// @Success 		200 {object} models.Product
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/product/create [POST]
func (h HandlerV1) CreateProduct(c *gin.Context) {
	var (
		body        models.CreateProduct
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Fatal("Create product error")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.ProductService().CreateProduct(ctx, &pbp.ProductWithCategoryId{
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
		h.Logger.Fatal("Create product error")
		return
	}

	c.JSON(http.StatusCreated, response)
}

// @Summary 		Get Product
// @Security 		ApiKeyAuth
// @Description 	Api for get product
// @Tags 			Product
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Id Product"
// @Success 		200 {object} models.Product
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/product/get/{id} [GET]
func (h *HandlerV1) GetProduct(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.ProductService().GetProduct(ctx, &pbp.GetProductRequest{
		Id: cast.ToInt64(id),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Fatal("Get product error")
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary 		Update Product
// @Security 		ApiKeyAuth
// @Description 	Api for update product
// @Tags 			Product
// @Accept 			json
// @Produce 		json
// @Param 			Product body models.UpdateProduct true "Update Product"
// @Success 		200 {object} models.Product
// @Failure 		400 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router	 		/v1/product/update [PUT]
func (h *HandlerV1) UpdateProduct(c *gin.Context) {
	var (
		body        models.UpdateProduct
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Fatal("Update product error")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.ProductService().UpdateProduct(ctx, &pbp.Product{
		Id:          body.Id,
		Title:       body.Title,
		Description: body.Description,
		Price:       body.Price,
		Discount:    body.Discount,
		Picture:     body.Picture,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Fatal("Update product error")
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary 		Delete Product
// @Security 		ApiKeyAuth
// @Description 	Api for delete product
// @Tags 			Product
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Id Product"
// @Success 		200 {object} models.CheckResponse
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/product/delete/{id} [DELETE]
func (h *HandlerV1) DeleteProduct(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.ProductService().DeleteProduct(ctx, &pbp.DeleteProductRequest{
		Id: cast.ToInt64(id),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Fatal("Delete product error")
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary	 		Get List Product
// @Security 		ApiKeyAuth
// @Description 	Api for get all product
// @Tags 			Product
// @Accept 			json
// @Produce 		json
// @Param 			page path string true "Page Product"
// @Param 			limit path string true "Limit Product"
// @Success 		200 {object} models.ProductList
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/products/get/{page}/{limit} [GET]
func (h *HandlerV1) ListProduct(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	page := c.Param("page")
	limit := c.Param("limit")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.ProductService().ListProduct(ctx, &pbp.GetAllRequest{
		Page:  cast.ToInt64(page),
		Limit: cast.ToInt64(limit),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Fatal("Get ListProduct error")
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary 		Search Product
// @Security 		ApiKeyAuth
// @Description 	Api for search product by title
// @Tags 			Product
// @Accept 			json
// @Produce 		json
// @Param 			Product body models.SearchProductRequest true "Search Products"
// @Success 		200 {object} models.ProductList
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/product/search [POST]
func (h *HandlerV1) SearchProduct(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
		body        models.SearchProductRequest
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Fatal("Search product error")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.ProductService().SearchProduct(ctx, &pbp.SearchProductRequest{
		Page:  body.Page,
		Limit: body.Limit,
		Title: body.Title,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Fatal("Search product error")
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary 		Get all product by category_id
// @Security 		ApiKeyAuth
// @Description 	Api for get all product by category_id
// @Tags 			Product
// @Accept 			json
// @Produce 		json
// @Param 			Product body models.GetAllProductByCategoryIdRequest true "Get Products"
// @Success 		200 {object} models.ProductList
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/products/get_category_id [POST]
func (h *HandlerV1) GetAllProductByCategoryId(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
		body        models.GetAllProductByCategoryIdRequest
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Fatal("Get products by category id error")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.ProductService().GetAllProductByCategoryId(ctx, &pbp.GetProductsByCategoryIdRequest{
		Id:    body.CategoryId,
		Page:  body.Page,
		Limit: body.Limit,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Fatal("Get products by category id error")
		return
	}

	c.JSON(http.StatusOK, response)
}
