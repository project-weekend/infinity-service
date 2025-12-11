package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/valkey-io/valkey-go"
	"gorm.io/gorm"

	"github.com/infinity/infinity-service/internal/entity"
	"github.com/infinity/infinity-service/internal/repository"
	"github.com/infinity/infinity-service/internal/repository/db"
)

const (
	cacheKeyProductCategory_All            = "product-category-repository:all"
	cacheKeyProductCategory_ByCategoryCode = "product-category-repository:code:%s"
	cacheTTL                               = 15 * time.Minute
)

type CacheProductCategoryRepository struct {
	Logger          *slog.Logger
	Cache           valkey.Client
	InnerRepository repository.ProductCategoryRepository
}

func NewCacheProductCategoryRepository(logger *slog.Logger, cache valkey.Client, innerRepository *db.MySqlProductCategoryRepository) CacheProductCategoryRepository {
	return CacheProductCategoryRepository{
		Logger:          logger,
		Cache:           cache,
		InnerRepository: innerRepository,
	}
}

func (r *CacheProductCategoryRepository) Save(ctx context.Context, db *gorm.DB, entity *entity.ProductCategory) error {
	// First, save to database via inner repository
	if err := r.InnerRepository.Save(ctx, db, entity); err != nil {
		return err
	}

	// Then, invalidate relevant caches
	helper := NewCacheHelper(r.Cache, r.Logger)
	cacheKeys := []string{
		cacheKeyProductCategory_All,
		fmt.Sprintf(cacheKeyProductCategory_ByCategoryCode, entity.CategoryCode),
	}

	for _, key := range cacheKeys {
		// Ignore cache delete errors to not fail the operation
		_ = helper.Delete(ctx, key)
	}

	return nil
}

func (r *CacheProductCategoryRepository) CountByCategoryCode(ctx context.Context, categoryCode string) (int64, error) {
	helper := NewCacheHelper(r.Cache, r.Logger)
	cacheKey := fmt.Sprintf(cacheKeyProductCategory_ByCategoryCode, categoryCode)

	cachedValue, err := helper.Get(ctx, cacheKey)
	if err == nil && cachedValue != "" {
		count, parseErr := strconv.ParseInt(cachedValue, 10, 64)
		if parseErr == nil {
			return count, nil
		}
	}

	count, err := r.InnerRepository.CountByCategoryCode(ctx, categoryCode)
	if err != nil {
		return 0, err
	}

	_ = helper.Set(ctx, cacheKey, strconv.FormatInt(count, 10), cacheTTL)

	return count, nil
}

func (r *CacheProductCategoryRepository) FindAll(ctx context.Context) ([]entity.ProductCategory, error) {
	helper := NewCacheHelper(r.Cache, r.Logger)
	cacheKey := cacheKeyProductCategory_All

	r.Logger.InfoContext(ctx, "FindAll: Attempting to get from cache", "key", cacheKey)

	// Try to get from cache
	cachedValue, err := helper.Get(ctx, cacheKey)
	if err == nil && cachedValue != "" {
		r.Logger.InfoContext(ctx, "FindAll: Cache HIT", "key", cacheKey, "valueLength", len(cachedValue))
		var categories []entity.ProductCategory
		if unmarshalErr := json.Unmarshal([]byte(cachedValue), &categories); unmarshalErr == nil {
			r.Logger.DebugContext(ctx, "FindAll: Successfully unmarshaled from cache", "count", len(categories))
			return categories, nil
		} else {
			r.Logger.WarnContext(ctx, "FindAll: Failed to unmarshal cached data", "error", unmarshalErr)
		}
	} else {
		r.Logger.InfoContext(ctx, "FindAll: Cache MISS", "key", cacheKey, "error", err)
	}

	r.Logger.DebugContext(ctx, "FindAll: Fetching from database")

	categories, err := r.InnerRepository.FindAll(ctx)
	if err != nil {
		r.Logger.ErrorContext(ctx, "FindAll: Database query failed", "error", err)
		return nil, err
	}

	r.Logger.InfoContext(ctx, "FindAll: Retrieved from database", "count", len(categories))

	if jsonData, marshalErr := json.Marshal(categories); marshalErr == nil {
		r.Logger.InfoContext(ctx, "FindAll: Storing in cache", "key", cacheKey, "dataLength", len(jsonData))
		if setErr := helper.Set(ctx, cacheKey, string(jsonData), cacheTTL); setErr != nil {
			r.Logger.ErrorContext(ctx, "FindAll: Failed to store in cache", "error", setErr)
		} else {
			r.Logger.DebugContext(ctx, "FindAll: Successfully stored in cache", "ttl", cacheTTL)
		}
	} else {
		r.Logger.ErrorContext(ctx, "FindAll: Failed to marshal categories", "error", marshalErr)
	}

	return categories, nil
}
