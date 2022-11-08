package cmd

import (
	"context"
	"embed"
	"github.com/go-redis/redis/v9"
	"real-estate/env"
	"real-estate/internal/cache"
	crawler2 "real-estate/internal/crawler"
	"real-estate/internal/database"
	"real-estate/server"
)

// Start the server
func Start(fs embed.FS) {
	// App context
	ctx := context.Background()

	// env config
	_env := env.GetEnv(".env.development")

	// Run database with env config
	//db-data := database.NewMySQLDatabase(ctx, _env).ConnectDB() // or work with mysql
	db := database.NewCockRoachDatabase(ctx, _env, fs).ConnectDB()
	defer db.Close()

	rdb := redis.NewClient(
		&redis.Options{
			Addr:     "redis:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		},
	)

	crawler := crawler2.New()
	defer crawler.MustClose()

	// Run server with context, database
	//server.NewGinServer(ctx, db-data, _env.SERVER_PORT).Run() // with Gin for example
	server.NewEchoServer(ctx, db, _env.SERVER_PORT, cache.New(rdb), crawler).Run()
}
