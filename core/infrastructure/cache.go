package estate

import (
	"github.com/go-redis/redis/v9"
	"real-estate/core/entities"
)

type Cache interface {
	Add(estate *entities.Estate)
	Get() *entities.Estate
}

type EstateCache struct {
	rdb *redis.Client

}