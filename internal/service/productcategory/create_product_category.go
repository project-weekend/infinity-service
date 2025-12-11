package productcategory

import (
	"context"
	"strings"

	"github.com/google/uuid"

	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/model"
)

func (p *ProductCategoryServiceImpl) Create(ctx context.Context, request *model.CreateProductCategoryRequest) (*model.GenericResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	categoryCode := generateProductCategoryCode(request.Name)
	total, err := p.ProductCategoryRepository.CountByCategoryCode(ctx, categoryCode)
	if err != nil {
		return nil, err
	}

	// regenerate category code if it already exists
	if total > 0 {
		categoryCode = generateProductCategoryCode(request.Name)
	}

	productCategory := &entity.ProductCategory{
		CategoryCode: generateProductCategoryCode(request.Name),
		Name:         request.Name,
		Description:  request.Description,
	}

	if err := p.ProductCategoryRepository.Save(ctx, productCategory); err != nil {
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &model.GenericResponse{
		Success: true,
	}, nil
}

func generateProductCategoryCode(name string) string {
	return strings.ToUpper(strings.ReplaceAll(name, " ", "_"))[:6] + "-" + uuid.New().String()[:6]
}
