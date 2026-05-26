package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	rdb *redis.Client
}

func New(rdb *redis.Client) *Cache {
	return &Cache{rdb: rdb}
}

// Set lưu value dưới dạng JSON
func (c *Cache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.rdb.Set(ctx, key, b, ttl).Err()
}

// Get lấy value và unmarshal vào dest
func (c *Cache) Get(ctx context.Context, key string, dest interface{}) error {
	val, err := c.rdb.Get(ctx, key).Bytes()
	if err != nil {
		return err // redis.Nil nếu không tìm thấy
	}
	return json.Unmarshal(val, dest)
}

// Delete xóa key
func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.rdb.Del(ctx, key).Err()
}

// Exists kiểm tra key có tồn tại không
func (c *Cache) Exists(ctx context.Context, key string) bool {
	n, _ := c.rdb.Exists(ctx, key).Result()
	return n > 0
}

// SetString lưu string thuần
func (c *Cache) SetString(ctx context.Context, key, value string, ttl time.Duration) error {
	return c.rdb.Set(ctx, key, value, ttl).Err()
}

// GetString lấy string thuần
func (c *Cache) GetString(ctx context.Context, key string) (string, error) {
	return c.rdb.Get(ctx, key).Result()
}

// IsNil kiểm tra lỗi có phải do key không tồn tại không
func IsNil(err error) bool {
	return err == redis.Nil
}
