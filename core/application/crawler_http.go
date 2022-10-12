package app_todo

import (
	"context"
	"database/sql"
	"github.com/go-redis/redis/v9"
	"github.com/go-rod/rod"
	"github.com/labstack/echo/v4"
	"net/http"
	"real-estate/core/infrastructure/cache"
	"real-estate/core/infrastructure/crawler"
	"real-estate/core/infrastructure/estate"
	"real-estate/core/infrastructure/storage"

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

	return c.JSON(http.StatusOK, estates)
}



// Constructor
func NewCrawlerHttpService(ctx context.Context,
	browser *rod.Browser,
	db *sql.DB,
	rdb *redis.Client,

	) *CrawlerHttpService {
	return &CrawlerHttpService{
		gtw: estate.NewLogic(
			crawler.NewOtodom(browser),
			storage.New(ctx, db),
			cache.New(),
			),
	}
}
