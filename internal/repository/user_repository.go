package repository

import (
	"github.com/infinity/infinity-service/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(db *gorm.DB, entity *entity.User, id string) error
}
