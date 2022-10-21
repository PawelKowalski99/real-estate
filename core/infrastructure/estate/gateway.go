package estate

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"real-estate/core/entities"
	"real-estate/core/infrastructure/crawler"
	//	"real-estate/core/infrastructure/storage"
	"real-estate/core/infrastructure/storage/estate"
	"real-estate/internal/cache"
)

// Gateway for access to EstateStorage, crawler and cache
type Gateway interface {
	GetEstates(mode, city, estateType string) ([]entities.Estate, error, int)
}

// Logic Domain
type Logic struct {
	crawler crawler.Crawler
	queries *estate.Queries
	cache   cache.Cache
}

// GetEstates ...
func (t *Logic) GetEstates(mode, city, estateType string) ([]entities.Estate, error, int) {
	ctx := context.Background()

	// Add parser

	cacheQuery := mode + city + estateType + "1"

	_, errCache := t.cache.Get(ctx, cacheQuery)

	var estates []entities.Estate

	switch errCache {

	case cache.NotExistError:
		var err error
		estates, err = t.crawler.GetEstates(mode, city, estateType, 1)
		if err != nil {
			return nil, err, http.StatusNotFound
		}
		errCache = t.cache.Set(ctx, cacheQuery)

		for _, est := range estates {
			_, err := t.queries.CreateEstate(ctx, estate.CreateEstateParams{
				ID:         uuid.New(),
				Urlstr:     est.URL,
				Addressstr: est.Address,
				Surface:    est.Surface,
				RoomAmount: est.RoomAmount,
				PricePerM2: est.PricePerM2,
				Price:      est.Price,
				Query:      mode + city + estateType,
			})
			if err != nil {
				return nil, nil, http.StatusBadRequest
			}
		}
		if err != nil {
			return nil, nil, http.StatusBadRequest
		}
		fmt.Println("Crawled estates count: ", len(estates))

	case nil:
		var err error

		estatesDb, err := t.queries.FindEstates(ctx, mode+city+estateType)
		if err != nil {
			return nil, err, http.StatusBadRequest
		}

		fmt.Println("Estates count from db: ", len(estatesDb))

		for _, est := range estatesDb {
			estates = append(estates, entities.Estate{
				URL:        est.Urlstr,
				Address:    est.Addressstr,
				Surface:    est.Surface,
				RoomAmount: est.RoomAmount,
				PricePerM2: est.PricePerM2,
				Price:      est.Price,
			})
		}

	default:
		return nil, errCache, http.StatusBadRequest
	}

	return estates, nil, http.StatusOK
}

// Constructor
func NewLogic(
	crawler crawler.Crawler,
	queries *estate.Queries,
	cache cache.Cache,

) *Logic {
	return &Logic{
		crawler: crawler,
		queries: queries,
		cache:   cache,
	}
}
