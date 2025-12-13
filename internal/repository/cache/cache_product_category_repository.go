package cache

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/repository"
	"github.com/infinity/infinity-service/server/config"
	"github.com/valkey-io/valkey-go"
	"gorm.io/gorm"
)

const (
	cacheKeyProductCategory_All  = "product-category:all"
	cacheKeyProductCategory_ByID = "product-category:id:%s"
)

type CacheProductCategoryRepository struct {
	AppConfig       *config.Config
	Cache           *Cache[entity.ProductCategory]
	InnerRepository repository.ProductCategoryRepository
}

func NewCacheProductCategoryRepository(logger *slog.Logger, appConfig *config.Config, client valkey.Client, innerRepository repository.ProductCategoryRepository) *CacheProductCategoryRepository {
	return &CacheProductCategoryRepository{
		AppConfig:       appConfig,
		Cache:           NewCache[entity.ProductCategory](logger, client),
		InnerRepository: innerRepository,
	}
}

func (r *CacheProductCategoryRepository) Create(ctx context.Context, db *gorm.DB, entity *entity.ProductCategory) error {
	if err := r.InnerRepository.Create(ctx, db, entity); err != nil {
		return err
	}
	// Invalidate list cache after create
	r.Cache.TryDeleteCache(ctx, cacheKeyProductCategory_All)
	return nil
}

func (r *CacheProductCategoryRepository) Delete(ctx context.Context, db *gorm.DB, entity *entity.ProductCategory) error {
	if err := r.InnerRepository.Delete(ctx, db, entity); err != nil {
		return err
	}
	// Invalidate caches after delete
	r.Cache.TryDeleteCache(ctx,
		cacheKeyProductCategory_All,
		fmt.Sprintf(cacheKeyProductCategory_ByID, fmt.Sprint(entity.ID)),
	)
	return nil
}

func (r *CacheProductCategoryRepository) CountByCategoryCode(ctx context.Context, db *gorm.DB, categoryCode string) (int64, error) {
	// No caching for count - delegate directly
	return r.InnerRepository.CountByCategoryCode(ctx, db, categoryCode)
}

func (r *CacheProductCategoryRepository) FindAll(ctx context.Context, db *gorm.DB) ([]entity.ProductCategory, error) {
	// Try load from cache
	if cached := r.Cache.TryLoadCache(ctx, cacheKeyProductCategory_All); cached != nil {
		return cached, nil
	}

	// Cache miss - fetch from DB
	entities, err := r.InnerRepository.FindAll(ctx, db)
	if err != nil {
		return nil, err
	}

	// Save to cache
	r.Cache.TrySaveCache(ctx, cacheKeyProductCategory_All, r.AppConfig.ValkeyConfig.TTLInMinutes, entities)
	return entities, nil
}

func (r *CacheProductCategoryRepository) FindByID(ctx context.Context, db *gorm.DB, category *entity.ProductCategory, id string) error {
	cacheKey := fmt.Sprintf(cacheKeyProductCategory_ByID, id)

	// Try load from cache
	if cached := r.Cache.TryLoadCache(ctx, cacheKey); len(cached) > 0 {
		*category = cached[0]
		return nil
	}

	// Cache miss - fetch from DB
	if err := r.InnerRepository.FindByID(ctx, db, category, id); err != nil {
		return err
	}

	// Save to cache (wrap single entity in slice for generic cache)
	r.Cache.TrySaveCache(ctx, cacheKey, r.AppConfig.ValkeyConfig.TTLInMinutes, []entity.ProductCategory{*category})
	return nil
}
