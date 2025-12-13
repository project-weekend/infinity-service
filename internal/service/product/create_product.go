package product

import (
	"context"
	"time"

	"github.com/infinity/infinity-service/internal/common"
	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/model"
)

func (p *ProductServiceImpl) Create(ctx context.Context, request *model.CreateProductRequest) error {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// Check if product already exists
	total, err := p.ProductRepository.CountByProductSKU(ctx, tx, request.ProductSKU)
	if err != nil {
		p.Logger.ErrorContext(ctx, "failed to count product by sku", "error", err)
		return common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	if total > 0 {
		p.Logger.WarnContext(ctx, "product already exists", "productSKU", request.ProductSKU)
		return common.NewServiceError(common.ErrCode_Forbidden, nil)
	}

	// Verify category exists
	total, err = p.ProductCategoryRepository.CountByCategoryCode(ctx, tx, request.CategoryCode)
	if err != nil {
		p.Logger.ErrorContext(ctx, "failed to count product category by code", "error", err)
		return common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}
	if total == 0 {
		p.Logger.WarnContext(ctx, "product category not found", "categoryCode", request.CategoryCode)
		return common.NewServiceError(common.ErrCode_Forbidden, nil)
	}

	product := &entity.Products{
		ProductSKU:   request.ProductSKU,
		CategoryCode: request.CategoryCode,
		Name:         request.Name,
		Description:  request.Description,
		Price:        request.Price,
		Quantity:     request.Quantity,
		Status:       string(entity.ProductStatus_Active),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := p.ProductRepository.Create(ctx, tx, product); err != nil {
		p.Logger.ErrorContext(ctx, "failed to create product", "error", err)
		return common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	if err := tx.Commit().Error; err != nil {
		p.Logger.ErrorContext(ctx, "transaction commit error", "error", err)
		return common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	return nil
}
