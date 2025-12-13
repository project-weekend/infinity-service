package db

import (
	"context"
	"log/slog"

	"github.com/infinity/infinity-service/internal/entity"
	"gorm.io/gorm"
)

type MySqlProductCategoryRepository struct {
	Repository[entity.ProductCategory]
	Logger *slog.Logger
}

func NewMySqlProductCategoryRepository(logger *slog.Logger) *MySqlProductCategoryRepository {
	return &MySqlProductCategoryRepository{
		Logger: logger,
	}
}

func (r *MySqlProductCategoryRepository) CountByCategoryCode(ctx context.Context, db *gorm.DB, categoryCode string) (int64, error) {
	var total int64
	err := db.WithContext(ctx).Model(&entity.ProductCategory{}).Where("category_code = ?", categoryCode).Count(&total).Error
	return total, err
}
