package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// HandleHealthCheck returns a simple OK response for probes and load balancers.
func HandleHealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"service": "legendary-succotash",
	})
}
