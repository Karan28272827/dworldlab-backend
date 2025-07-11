package videoconverter

import "github.com/gofiber/fiber/v2"

func Handler(c *fiber.Ctx) error {
	// TODO: implement VideoConverter logic
	return c.JSON(fiber.Map{"message": "VideoConverter executed"})
}
