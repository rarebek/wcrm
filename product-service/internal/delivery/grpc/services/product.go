package services

import (
	"context"
	"time"

	pbp "wcrm/product-service/genproto/product"
	"wcrm/product-service/internal/entity"
	"wcrm/product-service/internal/usecase"

	// grpcClient "wcrm/product-service/internal/infrastructure/grpc_service_clients"

	"go.uber.org/zap"
)

type productRPC struct {
	logger  *zap.Logger
	product usecase.Product
	// client  grpcClient.ServiceClients
	pbp.UnimplementedProductServiceServer
}

func UserRPC(logger *zap.Logger, productUsecase usecase.Product) pbp.ProductServiceServer {
	return &productRPC{
		logger:  logger,
		product: productUsecase,
	}
}

func (u productRPC) CreateProduct(ctx context.Context, product *pbp.Product) (*pbp.Product, error) {

	req_product := entity.Product{
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		Discount:    product.Discount,
		Picture:     product.Picture,
		CategoryId:  product.CategoryId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	res_product, err := u.product.CreateProduct(ctx, &req_product)

	if err != nil {
		u.logger.Error("Create product error", zap.Error(err))
		return &pbp.Product{}, nil
	}

	return &pbp.Product{
		Id:          res_product.Id,
		Title:       res_product.Title,
		Description: res_product.Description,
		Price:       res_product.Price,
		Discount:    res_product.Discount,
		Picture:     res_product.Picture,
		CategoryId:  res_product.CategoryId,
		CreatedAt:   res_product.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   res_product.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:   "",
	}, nil
}
func (u productRPC) GetProduct(ctx context.Context, id *pbp.GetProductRequest) (*pbp.Product, error) {

	reqMap := make(map[string]int64)
	reqMap["id"] = id.Id

	res, err := u.product.GetProduct(ctx, reqMap)

	if err != nil {
		u.logger.Error("get product error", zap.Error(err))
		return &pbp.Product{}, nil
	}

	return &pbp.Product{
		Id:          id.Id,
		Title:       res.Title,
		Description: res.Description,
		Price:       res.Price,
		Discount:    res.Discount,
		Picture:     res.Picture,
		CategoryId:  res.CategoryId,
		CreatedAt:   res.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   res.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:   "",
	}, nil
}
func (u productRPC) DeleteProduct(ctx context.Context, id *pbp.DeleteProductRequest) (*pbp.CheckResponse, error) {

	// fmt.Println("birinchi bu >>>>>>>>>>>>>>>")
	// pp.Println(id)

	reqMap := make(map[string]int64)
	reqMap["id"] = id.Id

	productReq, err := u.product.DeleteProduct(ctx, id.Id)

	if err != nil {
		u.logger.Error("delete product error", zap.Error(err))
		return &pbp.CheckResponse{Check: false}, nil
	}

	return &pbp.CheckResponse{
		Check: productReq.Check,
	}, nil
}

func (u productRPC) UpdateProduct(ctx context.Context, product *pbp.Product) (*pbp.Product, error) {

	// fmt.Println("birinchi bu >>>>>>>>>>>>>>>")
	// pp.Println(product)

	updated_product := entity.Product{
		Id:          product.Id,
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		Discount:    product.Discount,
		Picture:     product.Picture,
		CategoryId:  product.CategoryId,
		UpdatedAt:   time.Now(),
	}

	row, err := u.product.UpdateProduct(ctx, &updated_product)

	if err != nil {
		u.logger.Error("update product error", zap.Error(err))
		return nil, err
	}

	return &pbp.Product{
		Id:          row.Id,
		Title:       row.Title,
		Description: row.Description,
		Price:       row.Price,
		Discount:    row.Discount,
		Picture:     row.Picture,
		CategoryId:  row.CategoryId,
		CreatedAt:   row.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   row.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
func (u productRPC) ListProduct(ctx context.Context, req *pbp.GetAllRequest) (*pbp.GetAllResponse, error) {
	offset := req.Limit * (req.Page - 1)

	res_products, err := u.product.ListProduct(ctx, uint64(req.Limit), uint64(offset), map[string]string{})

	if err != nil {
		u.logger.Error("get all product error", zap.Error(err))
		return nil, err
	}

	var products pbp.GetAllResponse

	for _, in := range res_products {
		products.Products = append(products.Products, &pbp.Product{
			Id:          in.Id,
			Title:       in.Title,
			Description: in.Description,
			Price:       in.Price,
			Discount:    in.Discount,
			Picture:     in.Picture,
			CategoryId:  in.CategoryId,
			CreatedAt:   in.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   in.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &products, nil
}
