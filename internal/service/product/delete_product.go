package product

import (
	"context"

	"github.com/infinity/infinity-service/internal/common"
	"github.com/infinity/infinity-service/internal/model"
)

func (p *ProductServiceImpl) Delete(ctx context.Context, request *model.DeleteProductRequest) error {
	product, err := p.ProductRepository.FindByID(ctx, p.DB, request.ID)
	if err != nil {
		p.Logger.ErrorContext(ctx, "failed to find product by id", "error", err, "id", request.ID)
		return common.NewServiceError(common.ErrCode_ResourceNotFound, nil)
	}

	if err := p.ProductRepository.Delete(ctx, p.DB, product); err != nil {
		p.Logger.ErrorContext(ctx, "failed to delete product", "error", err, "id", request.ID)
		return common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	return nil
}


