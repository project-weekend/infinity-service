package repository

import (
	"context"

	"github.com/infinity/infinity-service/internal/entity"
	"gorm.io/gorm"
)

type RoleRepository interface {
	FindByID(ctx context.Context, db *gorm.DB, entity *entity.Role, id string) error
}
