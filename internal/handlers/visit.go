package handlers

import (
	"github.com/adriantombu/amazon-alternatives-api/internal/config"
	"github.com/adriantombu/amazon-alternatives-api/internal/entities"
	"github.com/adriantombu/amazon-alternatives-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

// PostVisits adds a new entry to the visit statistics
func PostVisits(c *fiber.Ctx) error {
	visit := entities.Visit{Asin: middlewares.PostVisitRequest.Asin, Country: middlewares.PostVisitRequest.Country}
	_ = config.DB.Create(&visit)

	product := entities.Product{Asin: middlewares.PostVisitRequest.Asin}
	_ = config.DB.FirstOrCreate(&product)

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
