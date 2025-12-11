package converter

import (
	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/model"
)

func ProductCategoryToResponse(productCategory *entity.ProductCategory) *model.ProductCategoryResponse {
	return &model.ProductCategoryResponse{
		ID:           productCategory.ID,
		CategoryCode: productCategory.CategoryCode,
		Name:         productCategory.Name,
		Description:  productCategory.Description,
		CreatedAt:    productCategory.CreatedAt,
		UpdatedAt:    productCategory.UpdatedAt,
	}
}
