package v1

import (
	_ "api-gateway/api/docs"
	"context"
	"net/http"
	"time"

	"api-gateway/api/models"
	pbp "api-gateway/genproto/product"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/k0kubun/pp"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary 		Create Category
// @Security 		ApiKeyAuth
// @Description 	Api for create category
// @Tags 			Category
// @Accept 			json
// @Produce 		json
// @Param 			Category body models.CreateCategory true "Create Category"
// @Success 		200 {object} models.Category
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/category/create [POST]
func (h HandlerV1) CreateCategory(c *gin.Context) {
	var (
		body        models.CreateCategory
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Fatal("Create Category error")
	}

	pp.Println("CATEGORY: ", body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.ProductService().CreateCategory(ctx, &pbp.Category{
		Id:      uuid.NewString(),
		OwnerId: body.OwnerId,
		Name:    body.Name,
		Image:   body.Image,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Fatal("Create Category error")
		return
	}

	c.JSON(http.StatusCreated, response)
}

// @Summary 		Get Category
// @Security 		ApiKeyAuth
// @Description 	Api for get Category
// @Tags 			Category
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Id Category"
// @Success 		200 {object} models.Category
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/category/get/{id} [GET]
func (h *HandlerV1) GetCategory(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.ProductService().GetCategory(ctx, &pbp.GetCategoryRequest{
		Id: id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Fatal("Get Category error")
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary 		Update Category
// @Security 		ApiKeyAuth
// @Description 	Api for update Category
// @Tags 			Category
// @Accept 			json
// @Produce 		json
// @Param 			Category body models.UpdateCategory true "Update Category"
// @Success 		200 {object} models.Category
// @Failure 		400 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router	 		/v1/category/update [PUT]
func (h *HandlerV1) UpdateCategory(c *gin.Context) {
	var (
		body        models.UpdateCategory
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Fatal("Update Category error")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.ProductService().UpdateCategory(ctx, &pbp.Category{
		Id:      body.Id,
		OwnerId: body.OwnerId,
		Name:    body.Name,
		Image:   body.Image,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Fatal("Update Category error")
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary 		Delete Category
// @Security 		ApiKeyAuth
// @Description 	Api for delete Category
// @Tags 			Category
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Id Category"
// @Success 		200 {object} models.CheckResponse
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/category/delete/{id} [DELETE]
func (h *HandlerV1) DeleteCategory(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.ProductService().DeleteCategory(ctx, &pbp.DeleteCategoryRequest{
		Id: id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Fatal("Delete Category error")
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary	 		Get List Category
// @Security 		ApiKeyAuth
// @Description 	Api for get all Category
// @Tags 			Category
// @Accept 			json
// @Produce 		json
// @Param 			Category body models.CategoryListRequset true "List Category"
// @Success 		200 {object} models.CategoryList
// @Failure 		404 {object} models.StandartError
// @Failure 		500 {object} models.StandartError
// @Router 			/v1/category/getall [POST]
func (h *HandlerV1) ListCategory(c *gin.Context) {
	var (
		body        models.CategoryListRequset
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Fatal("Update Category error")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	response, err := h.Service.ProductService().ListCategory(ctx, &pbp.GetAllRequest{
		Page:    body.Page,
		Limit:   body.Limit,
		OwnerId: body.OwnerId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Fatal("Get ListCategory error")
		return
	}

	c.JSON(http.StatusOK, response)
}
