package service

import (
	"context"

	"github.com/infinity/infinity-service/internal/model"
)

type IProductCategoryService interface {
	Create(ctx context.Context, request *model.CreateProductCategoryRequest) (*model.GenericResponse, error)
	List(ctx context.Context) ([]model.ProductCategoryResponse, error)
}
