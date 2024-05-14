package services

import (
	pbp "wcrm/product-service/genproto/product"
	"wcrm/product-service/internal/usecase"
	"go.uber.org/zap"
)

type UserRPC struct {
	logger   *zap.Logger
	product  usecase.Product
	category usecase.Category
	pbp.UnimplementedProductServiceServer
}

func NewRPC(logger *zap.Logger,
	productUsecase usecase.Product, 
	categoryUsecase usecase.Category) pbp.ProductServiceServer {
	return &UserRPC{
		logger:   logger,
		category: categoryUsecase,
		product:  productUsecase,
	}
}
