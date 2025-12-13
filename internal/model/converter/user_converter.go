package converter

import (
	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		UserCode:  user.UserCode,
		RoleID:    user.RoleID,
		Email:     user.Email,
		Status:    user.Status,
		CreatedBy: user.CreatedBy,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserToLoginResponse(user *entity.User, token string) *model.LoginResponse {
	return &model.LoginResponse{
		UserID: user.ID,
		Email:  user.Email,
		Token:  token,
	}
}
