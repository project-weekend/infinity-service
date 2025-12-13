package productcategory

import (
	"context"

	"github.com/infinity/infinity-service/internal/common"
	"github.com/infinity/infinity-service/internal/model"
	"github.com/infinity/infinity-service/internal/model/converter"
)

func (p *ProductCategoryServiceImpl) List(ctx context.Context) ([]model.ProductCategoryResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	productCategories, err := p.ProductCategoryRepository.FindAll(ctx, tx)
	if err != nil {
		p.Logger.ErrorContext(ctx, "failed to find product categories", "error", err)
		return nil, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	if err := tx.Commit().Error; err != nil {
		p.Logger.ErrorContext(ctx, "transaction commit error", "error", err)
		return nil, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	response := make([]model.ProductCategoryResponse, len(productCategories))
	for i, productCategory := range productCategories {
		response[i] = *converter.ProductCategoryToResponse(&productCategory)
	}

	return response, nil
}
