package db

import (
	"log/slog"

	"github.com/infinity/infinity-service/internal/entity"
	"gorm.io/gorm"
)

type MySqlRoleRepository struct {
	Repository[entity.Role]
	Logger *slog.Logger
}

func NewMySqlRoleRepository(logger *slog.Logger, db *gorm.DB) *MySqlRoleRepository {
	return &MySqlRoleRepository{
		Repository: Repository[entity.Role]{
			DB: db,
		},
		Logger: logger,
	}
}

func (r *MySqlRoleRepository) FindByID(db *gorm.DB, entity *entity.Role, id string) error {
	return db.Where("id = ?", id).First(entity).Error
}
