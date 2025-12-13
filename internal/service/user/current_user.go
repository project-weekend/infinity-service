package user

import (
	"context"

	"github.com/infinity/infinity-service/internal/common"
	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/model"
	"github.com/infinity/infinity-service/internal/model/converter"
)

func (u *UserServiceImpl) CurrentUser(ctx context.Context, id string) (*model.UserResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := new(entity.User)
	if err := u.UserRepository.FindByID(ctx, tx, user, id); err != nil {
		u.Logger.WarnContext(ctx, "Failed find user by id", "err", err)
		return nil, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	return converter.UserToResponse(user), nil
}
