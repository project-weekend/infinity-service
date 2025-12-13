package repository

import (
	"context"

	"github.com/infinity/infinity-service/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, db *gorm.DB, entity *entity.User) error
	CountByEmail(ctx context.Context, db *gorm.DB, email string) (int64, error)
	FindByEmail(ctx context.Context, db *gorm.DB, user *entity.User, email string) error
	FindByID(ctx context.Context, db *gorm.DB, user *entity.User, userId string) error
	FindByToken(ctx context.Context, db *gorm.DB, user *entity.User, token string) error
	FindAll(ctx context.Context, db *gorm.DB) ([]entity.User, error)
	DeleteByID(ctx context.Context, db *gorm.DB, id string) error
}
