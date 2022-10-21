package cache

import (
	"context"
	"github.com/go-redis/redis/v9"
	_ "github.com/lib/pq"
	"time"
)

type RedisCache struct {
	rdb *redis.Client
}

func (e *RedisCache) Set(ctx context.Context, key string) error {
	return e.rdb.Set(ctx, key, true, 24*time.Hour).Err()
}

func (e *RedisCache) Get(ctx context.Context, url string) (val string, err error) {
	val, errT := e.rdb.Get(ctx, url).Result()
	if errT == redis.Nil {
		err = NotExistError
	}
	return val, err
}

func New(rcl *redis.Client) *RedisCache {
	return &RedisCache{rdb: rcl}
}
