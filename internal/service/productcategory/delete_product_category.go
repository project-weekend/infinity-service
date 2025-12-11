package productcategory

import (
	"context"

	"github.com/infinity/infinity-service/internal/common"
	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/model"
)

func (p *ProductCategoryServiceImpl) DeleteProductCategory(ctx context.Context, request *model.DeleteProductCategoryRequest) (*model.GenericResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := new(entity.User)
	if err := p.UserRepository.FindByID(tx, user, request.UserID); err != nil {
		p.Logger.ErrorContext(ctx, "failed to find user by id", err)
		return nil, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	if entity.Role(user.Role.Name) != entity.Role_Admin {
		p.Logger.WarnContext(ctx, "User is not admin", "user", user.ID)
		return nil, common.NewServiceError(common.ErrCode_Forbidden, nil)
	}

	if err := p.ProductCategoryRepository.DeleteByID(ctx, request.ID); err != nil {
		p.Logger.ErrorContext(ctx, "failed to delete product category by id", err)
		return nil, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}
	
	return &model.GenericResponse{Success: true}, nil
}
