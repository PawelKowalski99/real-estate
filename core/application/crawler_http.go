package app_estate

import (
	"context"
	"database/sql"
	"github.com/go-rod/rod"
	"github.com/labstack/echo/v4"
	"real-estate/core/infrastructure/crawler"
	"real-estate/core/infrastructure/estate"
	sqlc "real-estate/core/infrastructure/storage/estate"
	//	"real-estate/core/infrastructure/storage"
	"real-estate/internal/cache"
	//"real-estate/internal/crawler"
)

// The HTTP Handler for TODO
type CrawlerHttpService struct {
	gtw estate.Gateway
}

func (t *CrawlerHttpService) GetEstates(c echo.Context) (err error) {
	estates, err, i := t.gtw.GetEstates(
		c.QueryParam("mode"),
		c.QueryParam("city"),
		c.QueryParam("estate_type"),
	)
	if err != nil {
		return c.JSON(i, err.Error())
	}

	return c.JSON(i, estates)
}

// Constructor
func NewCrawlerHttpService(ctx context.Context,
	browser *rod.Browser,
	db *sql.DB,
	cache cache.Cache,

) *CrawlerHttpService {
	return &CrawlerHttpService{
		gtw: estate.NewLogic(
			crawler.NewOtodom(browser),
			sqlc.New(db),
			cache,
		),
	}
}
