package usecase

import (
	"wcrm/product-service/internal/entity"
	"wcrm/product-service/internal/infrastructure/repository"

	"context"
	"time"
	"wcrm/product-service/internal/pkg/otlp"
)

const (
	serviceNameCategory = "categoryService"
	spanNameCategory    = "categoryUsecase"
)

type Category interface {
	CreateCategory(ctx context.Context, product *entity.Category) (*entity.Category, error)
	GetCategory(ctx context.Context, params map[string]int64) (*entity.Category, error)
	ListCategory(ctx context.Context, limit, offset uint64, filter map[string]string) (*entity.AllCategory, error)
	UpdateCategory(ctx context.Context, product *entity.Category) (*entity.Category, error)
	DeleteCategory(ctx context.Context, id int64) (*entity.CheckResponse, error)
}

type categoryService struct {
	BaseUseCase
	repo       repository.Category
	ctxTimeout time.Duration
	// client
}

func NewCategoryService(ctxTimeout time.Duration, repo repository.Category) categoryService {
	return categoryService{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (u categoryService) CreateCategory(ctx context.Context, category *entity.Category) (*entity.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameCategory, spanNameCategory+"Create")
	defer span.End()

	createdCategory, err := u.repo.CreateCategory(ctx, category)
	if err != nil {
		return &entity.Category{}, nil
	}

	return createdCategory, nil
}

func (u categoryService) GetCategory(ctx context.Context, params map[string]int64) (*entity.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameCategory, spanNameCategory+"Get")
	defer span.End()

	return u.repo.GetCategory(ctx, params)
}

func (u categoryService) ListCategory(ctx context.Context, limit, offset uint64, filter map[string]string) (*entity.AllCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameCategory, spanNameCategory+"List")
	defer span.End()

	return u.repo.ListCategory(ctx, limit, offset, filter)
}

func (u categoryService) UpdateCategory(ctx context.Context, product *entity.Category) (*entity.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameCategory, spanNameCategory+"Update")
	defer span.End()

	updatedCategory, err := u.repo.UpdateCategory(ctx, product)
	if err != nil {
		return &entity.Category{}, nil
	}

	return updatedCategory, nil
}

func (u categoryService) DeleteCategory(ctx context.Context, id int64) (*entity.CheckResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameCategory, spanNameCategory+"Delete")
	defer span.End()

	deleteCategory, err := u.repo.DeleteCategory(ctx, id)

	return deleteCategory, err
}
