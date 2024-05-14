package services

import (
	"context"
	"time"

	pbc "wcrm/product-service/genproto/product"
	"wcrm/product-service/internal/entity"

	"go.uber.org/zap"
)

func (u UserRPC) CreateCategory(ctx context.Context, category *pbc.Category) (*pbc.Category, error) {

	req_category := entity.Category{
		Name:      category.Name,
		Image:     category.Image,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	res_category, err := u.category.CreateCategory(ctx, &req_category)

	if err != nil {
		u.logger.Error("Create category error", zap.Error(err))
		return &pbc.Category{}, nil
	}

	return &pbc.Category{
		Id:        res_category.Id,
		Name:      res_category.Name,
		Image:     res_category.Image,
		CreatedAt: res_category.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: res_category.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
func (u UserRPC) GetCategory(ctx context.Context, id *pbc.GetCategoryRequest) (*pbc.Category, error) {

	reqMap := make(map[string]int64)
	reqMap["id"] = id.Id

	res, err := u.category.GetCategory(ctx, reqMap)

	if err != nil {
		u.logger.Error("get product error", zap.Error(err))
		return &pbc.Category{}, nil
	}

	return &pbc.Category{
		Id:        res.Id,
		Name:      res.Name,
		Image:     res.Image,
		CreatedAt: res.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: res.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
func (u UserRPC) DeleteCategory(ctx context.Context, id *pbc.DeleteCategoryRequest) (*pbc.CheckResponse, error) {

	categoryReq, err := u.category.DeleteCategory(ctx, id.Id)

	if err != nil {
		u.logger.Error("delete category error", zap.Error(err))
		return &pbc.CheckResponse{Check: false}, nil
	}

	return &pbc.CheckResponse{
		Check: categoryReq.Check,
	}, nil
}
func (u UserRPC) UpdateCategory(ctx context.Context, category *pbc.Category) (*pbc.Category, error) {

	updated_category := entity.Category{
		Id:        category.Id,
		Name:      category.Name,
		Image:     category.Image,
		UpdatedAt: time.Now(),
	}

	row, err := u.category.UpdateCategory(ctx, &updated_category)

	if err != nil {
		u.logger.Error("update category error", zap.Error(err))
		return nil, err
	}

	return &pbc.Category{
		Id:        row.Id,
		Name:      row.Name,
		Image:     row.Image,
		CreatedAt: row.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: row.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
func (u UserRPC) ListCategory(ctx context.Context, req *pbc.GetAllRequest) (*pbc.GetAllCategoryResponse, error) {
	offset := req.Limit * (req.Page - 1)

	res_categorys, err := u.category.ListCategory(ctx, uint64(req.Limit), uint64(offset), map[string]string{})

	if err != nil {
		u.logger.Error("get all category error", zap.Error(err))
		return nil, err
	}

	var category pbc.GetAllCategoryResponse

	for _, in := range res_categorys.Categories {
		category.Categories = append(category.Categories, &pbc.Category{
			Id:        in.Id,
			Name:      in.Name,
			Image:     in.Image,
			CreatedAt: in.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: in.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	category.Count = int64(res_categorys.Count)
	return &category, nil
}
