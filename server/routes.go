package server

import (
	"real-estate/core/application"
)

// a TODO Routes
func (es *EchoServer) estatesRoutes() {

	crawler := app_estate.NewCrawlerHttpService(es.ctx, es.crawler, es.db, es.cache)

	es.GET("/estate-crawler/estates", crawler.GetEstates)

}

// All routes
func (es *EchoServer) routes() {
	es.estatesRoutes()
}
