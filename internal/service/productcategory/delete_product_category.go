package productcategory

import (
	"context"

	"github.com/infinity/infinity-service/internal/common"
	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/model"
)

func (p *ProductCategoryServiceImpl) Delete(ctx context.Context, request *model.DeleteProductCategoryRequest) error {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	category := new(entity.ProductCategory)
	if err := p.ProductCategoryRepository.FindByID(ctx, p.DB, category, request.ID); err != nil {
		p.Logger.ErrorContext(ctx, "failed to find product category by id", "error", err)
		return common.NewServiceError(common.ErrCode_ResourceNotFound, nil)
	}

	if err := p.ProductCategoryRepository.Delete(ctx, p.DB, category); err != nil {
		p.Logger.ErrorContext(ctx, "failed to delete product category", "error", err)
		return common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	if err := tx.Commit().Error; err != nil {
		p.Logger.ErrorContext(ctx, "transaction commit error", "error", err)
		return common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	return nil
}
