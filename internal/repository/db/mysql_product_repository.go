package db

import (
	"context"
	"log/slog"

	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/model"
	"gorm.io/gorm"
)

type MySqlProductRepository struct {
	Repository[entity.Products]
	Logger *slog.Logger
}

func NewMySqlProductRepository(logger *slog.Logger) *MySqlProductRepository {
	return &MySqlProductRepository{
		Logger: logger,
	}
}

func (r *MySqlProductRepository) Create(ctx context.Context, db *gorm.DB, entity *entity.Products) error {
	return r.Repository.Create(ctx, db, entity)
}

func (r *MySqlProductRepository) CountByProductSKU(ctx context.Context, db *gorm.DB, productCode string) (int64, error) {
	var total int64
	err := db.WithContext(ctx).Model(&entity.Products{}).Where("product_sku = ?", productCode).Count(&total).Error
	return total, err
}

func (r *MySqlProductRepository) FindAll(ctx context.Context, db *gorm.DB) ([]entity.Products, error) {
	var products []entity.Products
	err := db.WithContext(ctx).Preload("Category").Find(&products).Error
	return products, err
}

func (r *MySqlProductRepository) FindByID(ctx context.Context, db *gorm.DB, id string) (*entity.Products, error) {
	var product entity.Products
	err := db.WithContext(ctx).Preload("Category").Where("id = ?", id).Take(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *MySqlProductRepository) Search(ctx context.Context, db *gorm.DB, request *model.SearchProductRequest) ([]entity.Products, int64, error) {
	var products []entity.Products
	if err := db.Preload("Category").Scopes(r.FilterContact(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.Products{}).Scopes(r.FilterContact(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (r *MySqlProductRepository) Delete(ctx context.Context, db *gorm.DB, entity *entity.Products) error {
	return r.Repository.Delete(ctx, db, entity)
}

func (r *MySqlProductRepository) FilterContact(request *model.SearchProductRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if productSKU := request.ProductSKU; productSKU != "" {
			productSKU = "%" + productSKU + "%"
			tx = tx.Where("product_sku LIKE ?", productSKU)
		}

		if name := request.Name; name != "" {
			name = "%" + name + "%"
			tx = tx.Where("name LIKE ?", name)
		}

		return tx
	}
}
