package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/valkey-io/valkey-go"
)

// CacheHelper provides convenient methods for common cache operations
type CacheHelper struct {
	Logger *slog.Logger
	Client valkey.Client
}

// NewCacheHelper creates a new cache helper instance
func NewCacheHelper(Client valkey.Client, logger *slog.Logger) *CacheHelper {
	return &CacheHelper{
		Client: Client,
		Logger: logger,
	}
}

// Set stores a value in cache with an expiration time
func (c *CacheHelper) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	var data string

	// Handle different value types
	switch v := value.(type) {
	case string:
		data = v
	case []byte:
		data = string(v)
	default:
		// Marshal to JSON for complex types
		jsonData, err := json.Marshal(value)
		if err != nil {
			return fmt.Errorf("failed to marshal value: %w", err)
		}
		data = string(jsonData)
	}

	cmd := c.Client.B().Set().Key(key).Value(data).ExSeconds(int64(expiration.Seconds())).Build()
	resp := c.Client.Do(ctx, cmd)

	return resp.Error()
}

// Get retrieves a value from cache
func (c *CacheHelper) Get(ctx context.Context, key string) (string, error) {
	cmd := c.Client.B().Get().Key(key).Build()
	resp := c.Client.Do(ctx, cmd)

	if resp.Error() != nil {
		return "", resp.Error()
	}

	return resp.ToString()
}

// GetJSON retrieves and unmarshals a JSON value from cache
func (c *CacheHelper) GetJSON(ctx context.Context, key string, dest interface{}) error {
	data, err := c.Get(ctx, key)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(data), dest)
}

// Delete removes a key from cache
func (c *CacheHelper) Delete(ctx context.Context, keys ...string) error {
	cmd := c.Client.B().Del().Key(keys...).Build()
	resp := c.Client.Do(ctx, cmd)

	return resp.Error()
}

// Exists checks if a key exists in cache
func (c *CacheHelper) Exists(ctx context.Context, keys ...string) (int64, error) {
	cmd := c.Client.B().Exists().Key(keys...).Build()
	resp := c.Client.Do(ctx, cmd)

	if resp.Error() != nil {
		return 0, resp.Error()
	}

	return resp.AsInt64()
}

// Expire sets an expiration time on a key
func (c *CacheHelper) Expire(ctx context.Context, key string, expiration time.Duration) error {
	cmd := c.Client.B().Expire().Key(key).Seconds(int64(expiration.Seconds())).Build()
	resp := c.Client.Do(ctx, cmd)

	return resp.Error()
}

// TTL returns the remaining time to live of a key
func (c *CacheHelper) TTL(ctx context.Context, key string) (time.Duration, error) {
	cmd := c.Client.B().Ttl().Key(key).Build()
	resp := c.Client.Do(ctx, cmd)

	if resp.Error() != nil {
		return 0, resp.Error()
	}

	seconds, err := resp.AsInt64()
	if err != nil {
		return 0, err
	}

	return time.Duration(seconds) * time.Second, nil
}

// Increment increments a numeric value in cache
func (c *CacheHelper) Increment(ctx context.Context, key string) (int64, error) {
	cmd := c.Client.B().Incr().Key(key).Build()
	resp := c.Client.Do(ctx, cmd)

	if resp.Error() != nil {
		return 0, resp.Error()
	}

	return resp.AsInt64()
}

// IncrementBy increments a numeric value by a specific amount
func (c *CacheHelper) IncrementBy(ctx context.Context, key string, value int64) (int64, error) {
	cmd := c.Client.B().Incrby().Key(key).Increment(value).Build()
	resp := c.Client.Do(ctx, cmd)

	if resp.Error() != nil {
		return 0, resp.Error()
	}

	return resp.AsInt64()
}

// Decrement decrements a numeric value in cache
func (c *CacheHelper) Decrement(ctx context.Context, key string) (int64, error) {
	cmd := c.Client.B().Decr().Key(key).Build()
	resp := c.Client.Do(ctx, cmd)

	if resp.Error() != nil {
		return 0, resp.Error()
	}

	return resp.AsInt64()
}

// DecrementBy decrements a numeric value by a specific amount
func (c *CacheHelper) DecrementBy(ctx context.Context, key string, value int64) (int64, error) {
	cmd := c.Client.B().Decrby().Key(key).Decrement(value).Build()
	resp := c.Client.Do(ctx, cmd)

	if resp.Error() != nil {
		return 0, resp.Error()
	}

	return resp.AsInt64()
}

// SetNX sets a value only if the key does not exist (useful for distributed locks)
func (c *CacheHelper) SetNX(ctx context.Context, key string, value string, expiration time.Duration) (bool, error) {
	cmd := c.Client.B().Setnx().Key(key).Value(value).Build()
	resp := c.Client.Do(ctx, cmd)

	if resp.Error() != nil {
		return false, resp.Error()
	}

	result, err := resp.AsInt64()
	if err != nil {
		return false, err
	}

	// If SetNX was successful, set the expiration
	if result == 1 {
		expCmd := c.Client.B().Expire().Key(key).Seconds(int64(expiration.Seconds())).Build()
		expResp := c.Client.Do(ctx, expCmd)
		if expResp.Error() != nil {
			return true, expResp.Error()
		}
	}

	return result == 1, nil
}

// MGet retrieves multiple values at once
func (c *CacheHelper) MGet(ctx context.Context, keys ...string) ([]string, error) {
	cmd := c.Client.B().Mget().Key(keys...).Build()
	resp := c.Client.Do(ctx, cmd)

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	values, err := resp.AsStrSlice()
	if err != nil {
		return nil, err
	}

	return values, nil
}

// MSet sets multiple key-value pairs at once
func (c *CacheHelper) MSet(ctx context.Context, pairs map[string]string) error {
	if len(pairs) == 0 {
		return nil
	}

	// Convert map to slice of alternating keys and values
	args := make([]string, 0, len(pairs)*2)
	for k, v := range pairs {
		args = append(args, k, v)
	}

	// Build command with first key-value pair
	if len(args) < 2 {
		return nil
	}

	builder := c.Client.B().Mset().KeyValue()
	for i := 0; i < len(args); i += 2 {
		if i == 0 {
			builder = builder.KeyValue(args[i], args[i+1])
		} else {
			builder = builder.KeyValue(args[i], args[i+1])
		}
	}

	cmd := builder.Build()
	resp := c.Client.Do(ctx, cmd)

	return resp.Error()
}

// FlushDB removes all keys from the current database (use with caution!)
func (c *CacheHelper) FlushDB(ctx context.Context) error {
	cmd := c.Client.B().Flushdb().Build()
	resp := c.Client.Do(ctx, cmd)

	return resp.Error()
}

// Ping tests the connection to the cache server
func (c *CacheHelper) Ping(ctx context.Context) error {
	cmd := c.Client.B().Ping().Build()
	resp := c.Client.Do(ctx, cmd)

	return resp.Error()
}

// Keys returns all keys matching a pattern
// WARNING: This command can be slow on large databases. Use with caution in production.
func (c *CacheHelper) Keys(ctx context.Context, pattern string) ([]string, error) {
	cmd := c.Client.B().Keys().Pattern(pattern).Build()
	resp := c.Client.Do(ctx, cmd)

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	return resp.AsStrSlice()
}

// HSet sets a field in a hash
func (c *CacheHelper) HSet(ctx context.Context, key, field, value string) error {
	cmd := c.Client.B().Hset().Key(key).FieldValue().FieldValue(field, value).Build()
	resp := c.Client.Do(ctx, cmd)

	return resp.Error()
}

// HGet gets a field from a hash
func (c *CacheHelper) HGet(ctx context.Context, key, field string) (string, error) {
	cmd := c.Client.B().Hget().Key(key).Field(field).Build()
	resp := c.Client.Do(ctx, cmd)

	if resp.Error() != nil {
		return "", resp.Error()
	}

	return resp.ToString()
}

// HGetAll gets all fields and values from a hash
func (c *CacheHelper) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	cmd := c.Client.B().Hgetall().Key(key).Build()
	resp := c.Client.Do(ctx, cmd)

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	return resp.AsStrMap()
}

// HDel deletes fields from a hash
func (c *CacheHelper) HDel(ctx context.Context, key string, fields ...string) error {
	cmd := c.Client.B().Hdel().Key(key).Field(fields...).Build()
	resp := c.Client.Do(ctx, cmd)

	return resp.Error()
}

// LPush prepends values to a list
func (c *CacheHelper) LPush(ctx context.Context, key string, values ...string) error {
	cmd := c.Client.B().Lpush().Key(key).Element(values...).Build()
	resp := c.Client.Do(ctx, cmd)

	return resp.Error()
}

// RPush appends values to a list
func (c *CacheHelper) RPush(ctx context.Context, key string, values ...string) error {
	cmd := c.Client.B().Rpush().Key(key).Element(values...).Build()
	resp := c.Client.Do(ctx, cmd)

	return resp.Error()
}

// LRange gets a range of elements from a list
func (c *CacheHelper) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	cmd := c.Client.B().Lrange().Key(key).Start(start).Stop(stop).Build()
	resp := c.Client.Do(ctx, cmd)

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	return resp.AsStrSlice()
}

// SAdd adds members to a set
func (c *CacheHelper) SAdd(ctx context.Context, key string, members ...string) error {
	cmd := c.Client.B().Sadd().Key(key).Member(members...).Build()
	resp := c.Client.Do(ctx, cmd)

	return resp.Error()
}

// SMembers returns all members of a set
func (c *CacheHelper) SMembers(ctx context.Context, key string) ([]string, error) {
	cmd := c.Client.B().Smembers().Key(key).Build()
	resp := c.Client.Do(ctx, cmd)

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	return resp.AsStrSlice()
}

// SIsMember checks if a value is a member of a set
func (c *CacheHelper) SIsMember(ctx context.Context, key, member string) (bool, error) {
	cmd := c.Client.B().Sismember().Key(key).Member(member).Build()
	resp := c.Client.Do(ctx, cmd)

	if resp.Error() != nil {
		return false, resp.Error()
	}

	result, err := resp.AsInt64()
	if err != nil {
		return false, err
	}

	return result == 1, nil
}

// ZAdd adds members to a sorted set with scores
func (c *CacheHelper) ZAdd(ctx context.Context, key string, score float64, member string) error {
	cmd := c.Client.B().Zadd().Key(key).ScoreMember().ScoreMember(score, member).Build()
	resp := c.Client.Do(ctx, cmd)

	return resp.Error()
}

// ZRange returns a range of members from a sorted set by index
func (c *CacheHelper) ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	cmd := c.Client.B().Zrange().Key(key).Min(fmt.Sprintf("%d", start)).Max(fmt.Sprintf("%d", stop)).Build()
	resp := c.Client.Do(ctx, cmd)

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	return resp.AsStrSlice()
}
