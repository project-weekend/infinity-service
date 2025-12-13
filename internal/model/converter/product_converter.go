package converter

import (
	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/model"
)

func ProductToResponse(product *entity.Products) *model.ProductResponse {
	return &model.ProductResponse{
		ID:          product.ID,
		ProductSKU:  product.ProductSKU,
		Category:    product.Category.Name,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Status:      product.Status,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
