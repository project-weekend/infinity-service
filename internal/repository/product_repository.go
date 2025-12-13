package repository

import (
	"context"

	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, db *gorm.DB, entity *entity.Products) error
	FindAll(ctx context.Context, db *gorm.DB) ([]entity.Products, error)
	FindByID(ctx context.Context, db *gorm.DB, id string) (*entity.Products, error)
	CountByProductSKU(ctx context.Context, db *gorm.DB, productSKU string) (int64, error)
	Search(ctx context.Context, db *gorm.DB, request *model.SearchProductRequest) ([]entity.Products, int64, error)
	Delete(ctx context.Context, db *gorm.DB, entity *entity.Products) error
}
