package service

import (
	"context"

	"github.com/infinity/infinity-service/internal/model"
)

type IUserService interface {
	Create(ctx context.Context, request *model.CreateUserRequest) (*model.UserResponse, error)
	Verify(ctx context.Context, request *model.VerifyUserRequest) (*model.Auth, error)
	Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error)
	CurrentUser(ctx context.Context, id string) (*model.UserResponse, error)
}
