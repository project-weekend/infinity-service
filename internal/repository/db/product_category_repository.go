package db

import (
	"context"
	"log/slog"

	"gorm.io/gorm"

	"github.com/infinity/infinity-service/internal/entity"
)

type MySqlProductCategoryRepository struct {
	Repository[entity.ProductCategory]
	Logger *slog.Logger
}

func NewMySqlProductCategoryRepository(logger *slog.Logger, db *gorm.DB) *MySqlProductCategoryRepository {
	return &MySqlProductCategoryRepository{
		Repository: Repository[entity.ProductCategory]{
			DB: db,
		},
		Logger: logger,
	}
}

func (r *MySqlProductCategoryRepository) Save(ctx context.Context, entity *entity.ProductCategory) error {
	return r.DB.WithContext(ctx).Save(entity).Error
}

func (r *MySqlProductCategoryRepository) CountByCategoryCode(ctx context.Context, categoryCode string) (int64, error) {
	var total int64
	err := r.DB.WithContext(ctx).Model(&entity.ProductCategory{}).Where("category_code = ?", categoryCode).Count(&total).Error
	return total, err
}

func (r *MySqlProductCategoryRepository) FindAll(ctx context.Context) ([]entity.ProductCategory, error) {
	var categories []entity.ProductCategory
	err := r.DB.WithContext(ctx).Find(&categories).Error
	return categories, err
}

func (r *MySqlProductCategoryRepository) FindByID(ctx context.Context, id string) (*entity.ProductCategory, error) {
	var category entity.ProductCategory
	err := r.DB.WithContext(ctx).First(&category, "id = ?", id).Error
	return &category, err
}

func (r *MySqlProductCategoryRepository) DeleteByID(ctx context.Context, id string) error {
	return r.DB.WithContext(ctx).Where("id = ?", id).Delete(&entity.ProductCategory{}).Error
}
