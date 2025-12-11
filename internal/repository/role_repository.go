package repository

import (
	"github.com/infinity/infinity-service/internal/entity"
	"gorm.io/gorm"
)

type RoleRepository interface {
	FindByID(db *gorm.DB, entity *entity.Role, id string) error
}
