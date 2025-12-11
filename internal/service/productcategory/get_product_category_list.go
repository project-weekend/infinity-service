package productcategory

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/infinity/infinity-service/internal/model"
	"github.com/infinity/infinity-service/internal/model/converter"
)

func (p *ProductCategoryServiceImpl) List(ctx context.Context) ([]model.ProductCategoryResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	productCategories, err := p.ProductCategoryRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		p.Logger.ErrorContext(ctx, "transaction commit error", err)
		return nil, fiber.ErrInternalServerError
	}

	response := make([]model.ProductCategoryResponse, len(productCategories))
	for i, productCategory := range productCategories {
		response[i] = *converter.ProductCategoryToResponse(&productCategory)
	}

	return response, nil
}
