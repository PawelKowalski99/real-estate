package server

import (
	"github.com/labstack/echo/v4/middleware"
	"real-estate/core/application"
)

// a TODO Routes
func (es *EchoServer) estatesRoutes() {

	crawler := app_estate.NewCrawlerHttpService(
		es.ctx, es.crawler, es.db, es.cache)

	g := es.Group("/estate-crawler")

	g.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(15)))

	g.GET("/estates", crawler.GetEstates)

	es.GET("/average-prices", crawler.GetAveragePrice)
	es.GET("/average-prices-per-m2", crawler.GetAveragePricePerM2)
}

// All routes
func (es *EchoServer) routes() {
	es.estatesRoutes()
}
