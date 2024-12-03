package main

import (
	"context"
	"log"
	"time"
	"urlShortener/config"
	"urlShortener/internal/handlers"
	mongodb "urlShortener/internal/repository/MongoDB"
	"urlShortener/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoURI).
		SetMaxPoolSize(5).
		SetMinPoolSize(1).
		SetMaxConnIdleTime(30*time.Second))

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	repo := mongodb.NewMongoRepository(client, cfg.DBName)
	generator := service.NewShortCodeGenerator(6)
	urlService := service.NewUrlService(repo, generator)

	handler := handlers.NewURLHandler(urlService, cfg.BaseUrl)

	e := echo.New()

	//Middleware

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//API
	e.Static("/", "static")
	e.POST("/api/shorten", handler.CreateShortURL)
	e.GET("/:code", handler.RedirectToURL)

	e.Start(":" + cfg.Port)

}
