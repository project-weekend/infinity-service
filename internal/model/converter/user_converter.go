package converter

import (
	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		UserID:    user.UserID,
		RoleID:    user.RoleID,
		Name:      user.Name,
		Email:     user.Email,
		Status:    string(user.Status),
		CreatedBy: user.CreatedBy,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserToLoginResponse(user *entity.User, token string) *model.LoginResponse {
	return &model.LoginResponse{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
		Token:  token,
	}
}
