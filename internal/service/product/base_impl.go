package product

import (
	"log/slog"

	"github.com/infinity/infinity-service/internal/repository"
	"github.com/infinity/infinity-service/server/config"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	Config                    *config.Config
	Logger                    *slog.Logger
	DB                        *gorm.DB
	UserRepository            repository.UserRepository
	ProductRepository         repository.ProductRepository
	ProductCategoryRepository repository.ProductCategoryRepository
}

func NewProductService(cfg *config.Config, logger *slog.Logger, db *gorm.DB, userRepository repository.UserRepository,
	productRepository repository.ProductRepository, productCategoryRepository repository.ProductCategoryRepository,
) *ProductServiceImpl {
	return &ProductServiceImpl{
		Config:                    cfg,
		Logger:                    logger,
		DB:                        db,
		UserRepository:            userRepository,
		ProductRepository:         productRepository,
		ProductCategoryRepository: productCategoryRepository,
	}
}
