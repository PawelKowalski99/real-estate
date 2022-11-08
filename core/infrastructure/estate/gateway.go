package estate

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"net/http"
	"real-estate/core/entities"
	"real-estate/core/infrastructure/crawler"
	"regexp"
	"strconv"
	"unicode"

	//	"real-estate/core/infrastructure/storage"
	"real-estate/core/infrastructure/storage/estate"
	"real-estate/internal/cache"
)

// Gateway for access to EstateStorage, crawler and cache
type Gateway interface {
	GetEstates(mode, city, estateType string) (int, error, int)
	GetAveragePrices(ctx context.Context) (map[string]AveragePrice, error, int)
	GetAveragePricesPerM2(ctx context.Context) (avgPrices map[string]AveragePrice, err error, httpErr int)
}

// Logic Domain
type Logic struct {
	crawler crawler.Crawler
	queries *estate.Queries
	cache   cache.Cache
}

// GetEstates ...
func (t *Logic) GetEstates(mode, city, estateType string) (int, error, int) {
	ctx := context.Background()

	buf := make(chan []entities.Estate, 30)
	ctx2, cancel := context.WithCancel(context.Background())

	var err error
	go func() {
		defer cancel()
		err = t.crawler.GetEstates(mode, city, estateType, buf)
		if err != nil {
			fmt.Println("GetEstates err: ", err.Error())
		}
	}()
	if err != nil {
		return 0, err, http.StatusNotFound
	}

	var estatesCount int

	for {
		fmt.Println("Length of estatesCrawl in buf:", len(buf))
		select {
		case estatesCrawl := <-buf:
			fmt.Println("Length of estatesCrawl", len(estatesCrawl))
			for _, est := range estatesCrawl {

				price, _, _ := transform.String(transform.Chain(runes.Remove(runes.In(unicode.Space)), runes.Map(func(r rune) rune {
					if unicode.IsPunct(r) {
						return '.'
					}
					return r
				})), est.Price)
				pricePerM2, _, _ := transform.String(transform.Chain(runes.Remove(runes.In(unicode.Space)), runes.Map(func(r rune) rune {
					if unicode.IsPunct(r) {
						return '.'
					}
					return r
				})), est.PricePerM2)

				var priceFloat, pricePerM2Float float64

				priceFloat, err = strconv.ParseFloat(price[:len(price)-3], 64)
				if err != nil {
					priceFloat = 0
					logrus.Error(err)
				}

				if pricePerM2 == "" {
					pricePerM2Float = 0
				} else {
					pricePerM2Float, err = strconv.ParseFloat(pricePerM2[:len(pricePerM2)-7], 64)
					if err != nil {
						pricePerM2Float = 0
						logrus.Error(err)
					}
				}

				idRegexp := regexp.MustCompile("([A-Z])\\w+")

				cacheQuery := mode + city + estateType + idRegexp.FindStringSubmatch(est.URL)[0]

				_, err3 := t.cache.Get(ctx, cacheQuery)
				if err3 == cache.NotExistError {
					surface, _ := strconv.ParseFloat(est.Surface[:len(est.Surface)-4], 64)

					_, err3 = t.queries.CreateEstate(ctx, estate.CreateEstateParams{

						ID:         uuid.New(),
						IDEstate:   idRegexp.FindStringSubmatch(est.URL)[0],
						Urlstr:     est.URL,
						Addressstr: est.Address,
						Surface:    surface,
						RoomAmount: est.RoomAmount,
						PricePerM2: pricePerM2Float,
						Price:      priceFloat,
						Query:      mode + city + estateType,
						City:       city,
						//RentPrice:  rentPrice,
					})
					if err3 != nil {
						fmt.Println(" AA ", err3.Error())
						return 0, err3, http.StatusBadRequest
					}
					err3 = t.cache.Set(ctx, cacheQuery)
					if err3 != nil {
						fmt.Println(" ", err3.Error())
						return 0, err3, http.StatusBadRequest
					}
				}
			}
			estatesCount++
			continue

		case <-ctx2.Done():
			break
		}
		break
	}

	close(buf)

	fmt.Println("Crawled pages of estates count: ", estatesCount)

	estatesDb, err := t.queries.FindEstates(ctx, mode+city+estateType)
	if err != nil {
		fmt.Println(" AA ", err.Error())
		return 0, err, http.StatusBadRequest
	}

	fmt.Println("Estates count from db: ", len(estatesDb))

	return len(estatesDb), nil, http.StatusOK
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
