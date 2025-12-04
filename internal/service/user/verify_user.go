package user

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/model"
)

func (u *UserServiceImpl) Verify(ctx context.Context, request *model.VerifyUserRequest) (*model.Auth, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// Hash the incoming raw token to match against the stored hash
	secretKey := u.Config.Security.SecretKey
	hashedToken := hashToken(request.Token, secretKey)

	session := new(entity.UserSession)
	if err := u.SessionRepository.FindByToken(tx, session, hashedToken); err != nil {
		u.Logger.WarnContext(ctx, "Failed find session by token", "err", err)
		return nil, fiber.ErrForbidden
	}

	if session.IsExpired() {
		u.Logger.WarnContext(ctx, "Session expired", "session_id", session.SessionID)
		return nil, fiber.ErrForbidden
	}

	if err := tx.Commit().Error; err != nil {
		u.Logger.WarnContext(ctx, "Failed commit transaction", "err", err)
		return nil, fiber.ErrInternalServerError
	}

	return &model.Auth{ID: session.UserID}, nil
}
