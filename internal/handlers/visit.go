package handlers

import (
	"github.com/adriantombu/amazon-alternatives-api/internal/config"
	"github.com/adriantombu/amazon-alternatives-api/internal/entities"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// PostVisits adds a new entry to the visit statistics
func PostVisits(c *fiber.Ctx) error {
	var visitRequest entities.PostVisitRequest

	if err := c.BodyParser(&visitRequest); err != nil {
		return c.JSON(fiber.Map{
			"status": "error",
			"data":   err.Error(),
		})
	}

	v := validator.New()
	if err := v.Struct(visitRequest); err != nil {
		return c.JSON(fiber.Map{
			"status": "error",
			"data":   err.Error(),
		})
	}

	visit := entities.Visit{Asin: visitRequest.Asin, Country: visitRequest.Country}
	_ = config.DB.Create(&visit)

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
