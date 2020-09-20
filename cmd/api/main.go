package main

import (
	"github.com/adriantombu/amazon-alternatives-api/internal/config"
	"github.com/adriantombu/amazon-alternatives-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config.LoadConfig()

	db, err := gorm.Open(postgres.Open(config.DatabaseUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	config.DB = db

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	app.Get("/", handlers.GetBase)
	app.Post("/visits", handlers.PostVisits)

	app.Listen(config.Port)
}
