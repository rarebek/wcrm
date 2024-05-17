package services

import (
	"context"
	"time"

	pbp "wcrm/product-service/genproto/product"
	"wcrm/product-service/internal/entity"

	"go.uber.org/zap"
)

func (u UserRPC) CreateProduct(ctx context.Context, product *pbp.ProductWithCategoryId) (*pbp.Product, error) {

	req_product := entity.ProductWithCategoryId{
		Id:          product.Id,
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		OwnerId:     product.OwnerId,
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
		OwnerId:     req_product.OwnerId,
		Title:       res_product.Title,
		Description: res_product.Description,
		Price:       res_product.Price,
		Discount:    res_product.Discount,
		Picture:     res_product.Picture,
		CreatedAt:   res_product.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   res_product.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:   "",
	}, nil
}
func (u UserRPC) GetProduct(ctx context.Context, id *pbp.GetProductRequest) (*pbp.Product, error) {

	reqMap := make(map[string]string)
	reqMap["id"] = id.Id

	res, err := u.product.GetProduct(ctx, reqMap)

	if err != nil {
		u.logger.Error("get product error", zap.Error(err))
		return &pbp.Product{}, nil
	}

	return &pbp.Product{
		Id:          res.Id,
		OwnerId:     res.OwnerId,
		Title:       res.Title,
		Description: res.Description,
		Price:       res.Price,
		Discount:    res.Discount,
		Picture:     res.Picture,
		CreatedAt:   res.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   res.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
func (u UserRPC) DeleteProduct(ctx context.Context, id *pbp.DeleteProductRequest) (*pbp.CheckResponse, error) {

	productReq, err := u.product.DeleteProduct(ctx, id.Id)

	if err != nil {
		u.logger.Error("delete product error", zap.Error(err))
		return &pbp.CheckResponse{Check: false}, nil
	}

	return &pbp.CheckResponse{
		Check: productReq.Check,
	}, nil
}
func (u UserRPC) UpdateProduct(ctx context.Context, product *pbp.Product) (*pbp.Product, error) {

	updated_product := entity.Product{
		Id:          product.Id,
		OwnerId:     product.OwnerId,
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		Discount:    product.Discount,
		Picture:     product.Picture,
		UpdatedAt:   time.Now(),
	}

	row, err := u.product.UpdateProduct(ctx, &updated_product)

	if err != nil {
		u.logger.Error("update product error", zap.Error(err))
		return nil, err
	}

	return &pbp.Product{
		Id:          row.Id,
		OwnerId:     row.OwnerId,
		Title:       row.Title,
		Description: row.Description,
		Price:       row.Price,
		Discount:    row.Discount,
		Picture:     row.Picture,
		CreatedAt:   row.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   row.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
func (u UserRPC) ListProduct(ctx context.Context, req *pbp.GetAllRequest) (*pbp.GetAllResponse, error) {
	offset := req.Limit * (req.Page - 1)

	filter := map[string]string{
		"owner_id": req.OwnerId,
	}

	res_products, err := u.product.ListProduct(ctx, uint64(req.Limit), uint64(offset), filter)

	if err != nil {
		u.logger.Error("get all product error", zap.Error(err))
		return nil, err
	}

	var products pbp.GetAllResponse

	for _, in := range res_products.Products {
		products.Products = append(products.Products, &pbp.Product{
			Id:          in.Id,
			OwnerId:     in.OwnerId,
			Title:       in.Title,
			Description: in.Description,
			Price:       in.Price,
			Discount:    in.Discount,
			Picture:     in.Picture,
			CreatedAt:   in.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   in.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	products.Count = int64(res_products.Count)
	return &products, nil
}
func (u UserRPC) SearchProduct(ctx context.Context, req *pbp.SearchProductRequest) (*pbp.SearchProductResponse, error) {

	offset := req.Limit * (req.Page - 1)

	res_products, err := u.product.SearchProduct(ctx, req.Limit, offset, req.Title, req.OwnerId)

	if err != nil {
		u.logger.Error("get products error", zap.Error(err))
		return nil, err
	}

	var products pbp.SearchProductResponse

	for _, in := range res_products.Products {
		products.Products = append(products.Products, &pbp.Product{
			Id:          in.Id,
			OwnerId:     in.OwnerId,
			Title:       in.Title,
			Description: in.Description,
			Price:       in.Price,
			Discount:    in.Discount,
			Picture:     in.Picture,
			CreatedAt:   in.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   in.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	products.Count = int64(res_products.Count)

	return &products, err
}
func (u UserRPC) GetAllProductByCategoryId(ctx context.Context, req *pbp.GetProductsByCategoryIdRequest) (*pbp.GetAllResponse, error) {

	offset := req.Limit * (req.Page - 1)

	res_products, err := u.product.GetAllProductByCategoryId(ctx, req.Limit, offset, req.CategoryId)

	if err != nil {
		u.logger.Error("get products error", zap.Error(err))
		return nil, err
	}

	var products pbp.GetAllResponse

	for _, in := range res_products.Products {
		products.Products = append(products.Products, &pbp.Product{
			Id:          in.Id,
			OwnerId:     in.OwnerId,
			Title:       in.Title,
			Description: in.Description,
			Price:       in.Price,
			Discount:    in.Discount,
			Picture:     in.Picture,
			CreatedAt:   in.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   in.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	products.Count = int64(res_products.Count)

	return &products, err
}
