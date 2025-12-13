package product

import (
	"context"

	"github.com/infinity/infinity-service/internal/common"
	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/model"
	"github.com/infinity/infinity-service/internal/model/converter"
)

func (p *ProductServiceImpl) Search(ctx context.Context, request *model.SearchProductRequest) ([]model.ProductResponse, int64, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := new(entity.User)
	if err := p.UserRepository.FindByID(ctx, tx, user, request.UserID); err != nil {
		p.Logger.ErrorContext(ctx, "failed to find user by id", "error", err)
		return nil, 0, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	products, total, err := p.ProductRepository.Search(ctx, tx, request)
	if err != nil {
		p.Logger.ErrorContext(ctx, "failed to search products", "error", err)
		return nil, 0, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	if err := tx.Commit().Error; err != nil {
		p.Logger.ErrorContext(ctx, "transaction commit error", "error", err)
		return nil, 0, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	responses := make([]model.ProductResponse, len(products))
	for i, product := range products {
		responses[i] = *converter.ProductToResponse(&product)
	}

	return responses, total, nil
}
