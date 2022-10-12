package server

import (
	app_todo "real-estate/core/application"
)

// a TODO Routes
func (es *EchoServer) toDoRoutes() {
	// call the TODO HTTP Service
	todo := app_todo.NewToDoHTTPService(es.ctx, es.db)

	crawler := app_todo.NewCrawlerHttpService(es.ctx, es.crawler, es.rdb)

	es.GET("/todo", todo.ListHandler)
	es.POST("/todo/create", todo.CreateHandler)

	es.GET("/estate-crawler/estates", crawler.GetEstates)


}

//func anotherRoutes()...

// All routes
func (es *EchoServer) routes() {
	es.toDoRoutes()
	//es.anotherRoutes()
}
