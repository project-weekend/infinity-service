package repository

import (
	"context"

	"github.com/infinity/infinity-service/internal/entity"
	"gorm.io/gorm"
)

type ProductCategoryRepository interface {
	CountByCategoryCode(ctx context.Context, categoryCode string) (int64, error)
	FindAll(ctx context.Context) ([]entity.ProductCategory, error)
	Save(ctx context.Context, db *gorm.DB, entity *entity.ProductCategory) error
}
