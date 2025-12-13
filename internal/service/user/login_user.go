package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/infinity/infinity-service/internal/common"
	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/model"
	"github.com/infinity/infinity-service/internal/model/converter"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (u *UserServiceImpl) Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	u.Logger.InfoContext(ctx, "user initiate login")

	user := new(entity.User)
	if err := u.UserRepository.FindByEmail(ctx, tx, user, request.Email); err != nil {
		u.Logger.InfoContext(ctx, "check email", "email", user.Email)
		if err == gorm.ErrRecordNotFound {
			u.Logger.WarnContext(ctx, "User not found", "email", request.Email)
			return nil, common.NewServiceError(common.ErrCode_Forbidden, nil)
		}
		u.Logger.WarnContext(ctx, "Failed to find user by email", "err", err)
		return nil, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	if user.Status != string(entity.UserStatus_Active) {
		u.Logger.WarnContext(ctx, "User is not active", "email", request.Email)
		return nil, common.NewServiceError(common.ErrCode_Forbidden, nil)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		u.Logger.WarnContext(ctx, "Invalid password", "email", request.Email)
		return nil, common.NewServiceError(common.ErrCode_Unauthorized, nil)
	}

	secretKey := u.Config.Security.SecretKey
	rawToken, hashedToken, err := generateSecureToken(secretKey)
	if err != nil {
		u.Logger.ErrorContext(ctx, "Failed to generate secure token", "err", err)
		return nil, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	session := &entity.UserSession{
		SessionCode: uuid.New().String(),
		UserID:      user.ID,
		Token:       hashedToken,
		ExpiresAt:   time.Now().Add(24 * time.Hour),
	}

	if err := u.SessionRepository.CreateSession(tx, session); err != nil {
		u.Logger.WarnContext(ctx, "Failed to create session", "err", err)
		return nil, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	if err := tx.Commit().Error; err != nil {
		u.Logger.WarnContext(ctx, "Failed commit transaction", "err", err)
		return nil, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	return converter.UserToLoginResponse(user, rawToken), nil
}
