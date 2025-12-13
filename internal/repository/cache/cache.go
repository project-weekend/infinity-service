package cache

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/valkey-io/valkey-go"
)

type Cache[T any] struct {
	Logger *slog.Logger
	Client valkey.Client
}

func NewCache[T any](logger *slog.Logger, client valkey.Client) *Cache[T] {
	return &Cache[T]{
		Logger: logger,
		Client: client,
	}
}

// TrySaveCache saves entities to cache with TTL in minutes. Logs errors but does not return them.
func (c *Cache[T]) TrySaveCache(ctx context.Context, cacheKey string, ttlMinutes int, entities []T) {
	dataBytes, err := json.Marshal(entities)
	if err != nil {
		c.Logger.ErrorContext(ctx, "failed to marshal entities", "key", cacheKey, "error", err)
		return
	}

	cmd := c.Client.B().Set().Key(cacheKey).Value(string(dataBytes)).ExSeconds(int64(ttlMinutes * 60)).Build()
	if err := c.Client.Do(ctx, cmd).Error(); err != nil {
		c.Logger.ErrorContext(ctx, "failed to save cache", "key", cacheKey, "error", err)
	}
}

// TryLoadCache attempts to load entities from cache. Returns nil if cache miss or error.
func (c *Cache[T]) TryLoadCache(ctx context.Context, cacheKey string) []T {
	cmd := c.Client.B().Get().Key(cacheKey).Build()
	resp := c.Client.Do(ctx, cmd)

	if resp.Error() != nil {
		if !valkey.IsValkeyNil(resp.Error()) {
			c.Logger.ErrorContext(ctx, "failed to get cache", "key", cacheKey, "error", resp.Error())
		}
		return nil
	}

	data, err := resp.ToString()
	if err != nil {
		c.Logger.ErrorContext(ctx, "failed to read cache response", "key", cacheKey, "error", err)
		return nil
	}

	var entities []T
	if err := json.Unmarshal([]byte(data), &entities); err != nil {
		c.Logger.ErrorContext(ctx, "failed to unmarshal cache data", "key", cacheKey, "error", err)
		return nil
	}

	return entities
}

// TryDeleteCache deletes cache keys. Logs errors but does not return them.
func (c *Cache[T]) TryDeleteCache(ctx context.Context, cacheKeys ...string) {
	if len(cacheKeys) == 0 {
		return
	}

	cmd := c.Client.B().Del().Key(cacheKeys...).Build()
	if err := c.Client.Do(ctx, cmd).Error(); err != nil {
		c.Logger.ErrorContext(ctx, "failed to delete cache", "keys", cacheKeys, "error", err)
	}
}
