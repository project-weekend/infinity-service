package repository

import (
	"context"

	"github.com/infinity/infinity-service/internal/entity"
	"gorm.io/gorm"
)

type ProductCategoryRepository interface {
	Create(ctx context.Context, db *gorm.DB, entity *entity.ProductCategory) error
	Delete(ctx context.Context, db *gorm.DB, entity *entity.ProductCategory) error
	CountByCategoryCode(ctx context.Context, db *gorm.DB, categoryCode string) (int64, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]entity.ProductCategory, error)
	FindByID(ctx context.Context, db *gorm.DB, entity *entity.ProductCategory, id string) error
}
