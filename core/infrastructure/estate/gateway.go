package estate

import (
	"context"
	"github.com/go-redis/redis/v9"
	"net/http"
	"real-estate/core/entities"
	"real-estate/core/infrastructure/cache"
	"real-estate/core/infrastructure/crawler"
	"real-estate/core/infrastructure/storage"
	"time"
)

// Gateway for access to EstateStorage, crawler and cache
type Gateway interface {
	GetEstates(mode, city, estateType string) ([]entities.Estate, error, int)
}

// Logic Domain
type Logic struct {
	crawler crawler.Crawler
	db      storage.EstateStorage
	cache     cache.Cache
}

// GetEstates ...
func (t *Logic) GetEstates(mode, city, estateType string) ([]entities.Estate, error, int) {
	ctx := context.Background()

	// Add parser

	queried := t.cache.Get(ctx,
		time.Now().String()+
		mode+city+estateType,
	)

	var estates []entities.Estate

	switch queried {

	case redis.Nil:
		var err error
		estates, err = t.crawler.GetEstates(mode, city, estateType, 1)
		if err != nil {
			return nil, err, http.StatusNotFound
		}
		_ = t.cache.Add(ctx, time.Now().String()+mode+city+estateType)


		go func() {
			goErr := func() error {
				err = t.db.CreateEstates(estates)
				if err != nil {

					return err
				}
				return nil
			}()
			if goErr != nil {
			}
		}()

	case nil:
		var err error
		estates, err = t.db.GetEstates(mode, city, estateType)
		if err != nil {
			return nil, err, http.StatusBadRequest
		}

	default:
		return nil, queried, http.StatusBadRequest
	}

	return estates, nil, 0
}

// Constructor
func NewLogic (
	crawler crawler.Crawler,
	estateStorage storage.EstateStorage,
	cache cache.Cache,
) *Logic {
	return &Logic{
		crawler:   crawler,
		db:        estateStorage,
		cache:     cache,
	}
}
