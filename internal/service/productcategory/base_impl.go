package productcategory

import (
	"log/slog"

	"github.com/infinity/infinity-service/internal/repository"
	"github.com/infinity/infinity-service/server/config"
	"gorm.io/gorm"
)

type ProductCategoryServiceImpl struct {
	Config                    *config.Config
	Logger                    *slog.Logger
	DB                        *gorm.DB
	ProductCategoryRepository repository.ProductCategoryRepository
	UserRepository            repository.UserRepository
	RoleRepository            repository.RoleRepository
}

func NewProductCategoryService(cfg *config.Config, logger *slog.Logger, db *gorm.DB, productCategoryRepository repository.ProductCategoryRepository) *ProductCategoryServiceImpl {
	return &ProductCategoryServiceImpl{
		Config:                    cfg,
		Logger:                    logger,
		DB:                        db,
		ProductCategoryRepository: productCategoryRepository,
	}
}
