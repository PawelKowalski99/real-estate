package cache

import (
	"context"
	"github.com/go-redis/redis/v9"
	"time"
)

type Cache interface {
	Add(ctx context.Context, url string) error
	Get(ctx context.Context, url string) error
}

type EstateCache struct {
	rdb *redis.Client
}

func (e EstateCache) Add(ctx context.Context, url string) error {

	e.rdb.Set(ctx, "", true, time.Hour)
	return nil
}

func (e EstateCache) Get(ctx context.Context, url string) error {
	return nil
}

func New() *EstateCache {
	return nil
}