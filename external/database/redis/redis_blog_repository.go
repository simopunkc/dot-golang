package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type BlogCache struct {
	rdb *redis.Client
}

func NewBlogCache(rdb *redis.Client) *BlogCache {
	return &BlogCache{rdb}
}

func (bc BlogCache) Get(key string) (string, error) {
	result, err := bc.rdb.Get(ctx, key).Result()
	return result, err
}

func (bc BlogCache) Set(key string, val string) error {
	err := bc.rdb.Set(ctx, key, val, 0).Err()
	return err
}

func (bc BlogCache) Del(key string) error {
	err := bc.rdb.Del(ctx, key).Err()
	return err
}

func (bc BlogCache) Exists(key string) bool {
	result, err := bc.rdb.Keys(ctx, key).Result()
	return len(result) > 0 && err == nil
}

func (bc BlogCache) Keys(prefixKey string) ([]string, error) {
	result, err := bc.rdb.Keys(ctx, prefixKey).Result()
	return result, err
}
