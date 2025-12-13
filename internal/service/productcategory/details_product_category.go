package productcategory

import (
	"context"

	"github.com/infinity/infinity-service/internal/common"
	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/model"
	"github.com/infinity/infinity-service/internal/model/converter"
)

func (p *ProductCategoryServiceImpl) Get(ctx context.Context, request *model.GetProductCategoryRequest) (*model.ProductCategoryResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	category := new(entity.ProductCategory)
	if err := p.ProductCategoryRepository.FindByID(ctx, tx, category, request.ID); err != nil {
		p.Logger.ErrorContext(ctx, "failed to find product category by id", "error", err)
		return nil, common.NewServiceError(common.ErrCode_ResourceNotFound, nil)
	}

	if err := tx.Commit().Error; err != nil {
		p.Logger.ErrorContext(ctx, "transaction commit error", "error", err)
		return nil, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	return converter.ProductCategoryToResponse(category), nil
}
