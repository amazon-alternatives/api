package handlers

import "github.com/gofiber/fiber/v2"

// GetBase is the health check endpoint
func GetBase(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": "Hello, World 👋!",
	})
}
