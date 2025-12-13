package service

import (
	"context"

	"github.com/infinity/infinity-service/internal/model"
)

type IProductService interface {
	Create(ctx context.Context, request *model.CreateProductRequest) error
	Search(ctx context.Context, request *model.SearchProductRequest) ([]model.ProductResponse, int64, error)
	Get(ctx context.Context, request *model.GetProductRequest) (*model.ProductResponse, error)
	Delete(ctx context.Context, request *model.DeleteProductRequest) error
}
