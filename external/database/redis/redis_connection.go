package redis

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func NewConfigRedis(dsn string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: "",
		DB:       0,
	})
	return rdb
}

func NewConnectionRedis() *redis.Client {
	dsn := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	return NewConfigRedis(dsn)
}
