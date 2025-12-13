package user

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/infinity/infinity-service/internal/common"
	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/model"
	"github.com/infinity/infinity-service/internal/model/converter"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserServiceImpl) Create(ctx context.Context, request *model.CreateUserRequest) (*model.UserResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	auth, ok := model.GetAuthFromContext(ctx)
	if !ok {
		u.Logger.WarnContext(ctx, "No authenticated user in context")
		return nil, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	currentUserID := auth.ID
	loggedInUser := new(entity.User)
	if err := u.UserRepository.FindByID(ctx, tx, loggedInUser, currentUserID); err != nil {
		u.Logger.WarnContext(ctx, "Failed find user by id", "err", err)
		return nil, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	if entity.Role(loggedInUser.Role.Name) != entity.Role_Admin {
		u.Logger.WarnContext(ctx, "User is not admin", "user", loggedInUser.ID)
		return nil, common.NewServiceError(common.ErrCode_Forbidden, []common.ErrorDetail{
			{
				ErrorCode: "FORBIDDEN",
				Message:   "Only admin users can create new users",
			},
		})
	}

	total, err := u.UserRepository.CountByEmail(ctx, tx, request.Email)
	if err != nil {
		u.Logger.WarnContext(ctx, "Failed count user by email", "err", err)
		return nil, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}
	if total > 0 {
		u.Logger.WarnContext(ctx, "User email already exists", "email", request.Email)
		return nil, common.NewServiceError(common.ErrCode_Forbidden, nil)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		u.Logger.WarnContext(ctx, "Failed generate password", "err", err)
		return nil, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	user := &entity.User{
		Email:     request.Email,
		RoleID:    request.RoleID,
		UserCode:  uuid.New().String(),
		Status:    string(entity.UserStatus_Active),
		Password:  string(password),
		CreatedBy: currentUserID,
	}

	if err := u.UserRepository.Save(ctx, tx, user); err != nil {
		u.Logger.WarnContext(ctx, "Failed save user", "err", err)
		return nil, common.NewServiceError(common.ErrCode_InternalServerError, nil)
	}

	if err := tx.Commit().Error; err != nil {
		u.Logger.WarnContext(ctx, "Failed commit transaction", "err", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToResponse(user), nil
}
