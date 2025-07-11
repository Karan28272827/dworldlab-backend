package jsonformatter

import (
	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx) error {
	mode := c.Query("mode", "pretty")
	data := c.Query("data", "")

	formatted, err := FormatJSON(data, mode)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}
	return c.JSON(fiber.Map{
		"mode":     mode,
		"formatted": formatted,
	})
}
