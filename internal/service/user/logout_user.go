package user

import (
	"context"

	"github.com/infinity/infinity-service/internal/common"
	"github.com/infinity/infinity-service/internal/model"
)

func (u *UserServiceImpl) Logout(ctx context.Context, request *model.LogoutRequest) error {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := u.SessionRepository.DeleteSession(tx, request.Token); err != nil {
		u.Logger.WarnContext(ctx, "Failed to delete session", "err", err)
		return common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	if err := tx.Commit().Error; err != nil {
		u.Logger.WarnContext(ctx, "Failed commit transaction", "err", err)
		return common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	return nil
}
