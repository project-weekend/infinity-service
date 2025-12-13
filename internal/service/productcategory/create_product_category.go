package productcategory

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/infinity/infinity-service/internal/common"
	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/model"
)

func (p *ProductCategoryServiceImpl) Create(ctx context.Context, request *model.CreateProductCategoryRequest) (*model.GenericResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	categoryCode := generateProductCategoryCode(request.Name)
	total, err := p.ProductCategoryRepository.CountByCategoryCode(ctx, tx, categoryCode)
	if err != nil {
		p.Logger.ErrorContext(ctx, "failed to count product category by code", "error", err)
		return nil, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	// regenerate category code if it already exists
	if total > 0 {
		p.Logger.WarnContext(ctx, "product category code already exists", "categoryCode", categoryCode)
		categoryCode = generateProductCategoryCode(request.Name)
	}

	productCategory := &entity.ProductCategory{
		CategoryCode: categoryCode,
		Name:         request.Name,
		Description:  request.Description,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := p.ProductCategoryRepository.Create(ctx, tx, productCategory); err != nil {
		p.Logger.ErrorContext(ctx, "failed to create product category", "error", err)
		return nil, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	if err := tx.Commit().Error; err != nil {
		p.Logger.ErrorContext(ctx, "transaction commit error", "error", err)
		return nil, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	return &model.GenericResponse{
		Success: true,
	}, nil
}

func generateProductCategoryCode(name string) string {
	return strings.ToUpper(strings.ReplaceAll(name, " ", "_"))[:6] + "-" + uuid.New().String()[:6]
}
