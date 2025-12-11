package productcategory

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/infinity/infinity-service/internal/model"
	"github.com/infinity/infinity-service/internal/model/converter"
)

func (p *ProductCategoryServiceImpl) Get(ctx context.Context, request *model.GetProductCategoryRequest) (*model.ProductCategoryResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	productCategories, err := p.ProductCategoryRepository.FindByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		p.Logger.ErrorContext(ctx, "transaction commit error", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ProductCategoryToResponse(productCategories), nil
}
