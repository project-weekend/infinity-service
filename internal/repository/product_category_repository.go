package repository

import (
	"context"

	"github.com/infinity/infinity-service/internal/entity"
)

type ProductCategoryRepository interface {
	CountByCategoryCode(ctx context.Context, categoryCode string) (int64, error)
	FindAll(ctx context.Context) ([]entity.ProductCategory, error)
	FindByID(ctx context.Context, id string) (*entity.ProductCategory, error)
	Save(ctx context.Context, entity *entity.ProductCategory) error
	DeleteByID(ctx context.Context, id string) error
}
