package user

import (
	"log/slog"

	repository "github.com/infinity/infinity-service/internal/repository/db"
	"github.com/infinity/infinity-service/server/config"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	Config            *config.Config
	Logger            *slog.Logger
	DB                *gorm.DB
	UserRepository    *repository.UserRepository
	SessionRepository *repository.SessionRepository
}

func NewUserService(config *config.Config, logger *slog.Logger, db *gorm.DB,
	userRepository *repository.UserRepository, sessionRepository *repository.SessionRepository,
) *UserServiceImpl {
	return &UserServiceImpl{
		Config:            config,
		Logger:            logger,
		DB:                db,
		UserRepository:    userRepository,
		SessionRepository: sessionRepository,
	}
}
