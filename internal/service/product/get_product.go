package product

import (
	"context"

	"github.com/infinity/infinity-service/internal/common"
	"github.com/infinity/infinity-service/internal/model"
	"github.com/infinity/infinity-service/internal/model/converter"
)

func (p *ProductServiceImpl) Get(ctx context.Context, request *model.GetProductRequest) (*model.ProductResponse, error) {
	product, err := p.ProductRepository.FindByID(ctx, p.DB, request.ID)
	if err != nil {
		p.Logger.ErrorContext(ctx, "failed to find product by id", "error", err, "id", request.ID)
		return nil, common.NewServiceError(common.ErrCode_ResourceNotFound, nil)
	}

	response := converter.ProductToResponse(product)
	return response, nil
}
