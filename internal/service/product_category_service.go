package service

import (
	"context"

	"github.com/infinity/infinity-service/internal/model"
)

type IProductCategoryService interface {
	Create(ctx context.Context, request *model.CreateProductCategoryRequest) (*model.GenericResponse, error)
	List(ctx context.Context) ([]model.ProductCategoryResponse, error)
	Get(ctx context.Context, request *model.GetProductCategoryRequest) (*model.ProductCategoryResponse, error)
	Delete(ctx context.Context, request *model.DeleteProductCategoryRequest) error
}
