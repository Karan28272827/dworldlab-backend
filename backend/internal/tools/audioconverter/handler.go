package audioconverter

import "github.com/gofiber/fiber/v2"

func Handler(c *fiber.Ctx) error {
	// TODO: implement AudioConverter logic
	return c.JSON(fiber.Map{"message": "AudioConverter executed"})
}
