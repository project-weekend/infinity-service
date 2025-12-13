package db

import (
	"context"
	"log/slog"

	"github.com/infinity/infinity-service/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	Logger *slog.Logger
}

func NewUserRepository(logger *slog.Logger) *UserRepository {
	return &UserRepository{
		Logger: logger,
	}
}

func (r *UserRepository) Save(ctx context.Context, db *gorm.DB, entity *entity.User) error {
	return db.WithContext(ctx).Save(entity).Error
}

func (r *UserRepository) CountByEmail(ctx context.Context, db *gorm.DB, email string) (int64, error) {
	var total int64
	err := db.Model(new(entity.User)).Where("email = ?", email).Count(&total).Error
	return total, err
}

func (r *UserRepository) FindByEmail(ctx context.Context, db *gorm.DB, user *entity.User, email string) error {
	return db.Where("email = ?", email).First(user).Error
}

func (r *UserRepository) FindByID(ctx context.Context, db *gorm.DB, user *entity.User, id string) error {
	return db.Preload("Role").Where("id = ?", id).First(user).Error
}

func (r *UserRepository) FindByToken(ctx context.Context, db *gorm.DB, user *entity.User, token string) error {
	return db.Where("token = ?", token).First(user).Error
}

func (r *UserRepository) FindAll(ctx context.Context, db *gorm.DB) ([]entity.User, error) {
	var users []entity.User
	err := db.Preload("Role").Find(&users).Error
	return users, err
}

func (r *UserRepository) DeleteByID(ctx context.Context, db *gorm.DB, id string) error {
	return db.Where("id = ?", id).Delete(&entity.User{}).Error
}
