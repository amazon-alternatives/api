package middlewares

import (
	"github.com/adriantombu/amazon-alternatives-api/internal/entities"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// PostVisitRequest represents the body of POST /visit
var PostVisitRequest entities.PostVisitRequest

// PostVisitsParser checks if the body can be parsed
func PostVisitsParser(c *fiber.Ctx) error {
	if err := c.BodyParser(&PostVisitRequest); err != nil {
		return c.JSON(fiber.Map{
			"status": "error",
			"data":   err.Error(),
		})
	}

	return c.Next()
}

// PostVisitsValidator checks if the body is valid
func PostVisitsValidator(c *fiber.Ctx) error {
	v := validator.New()
	if err := v.Struct(PostVisitRequest); err != nil {
		return c.JSON(fiber.Map{
			"status": "error",
			"data":   err.Error(),
		})
	}

	return c.Next()
}
