package audiotrimmer

import "github.com/gofiber/fiber/v2"

func Handler(c *fiber.Ctx) error {
	// TODO: implement AudioTrimmer logic
	return c.JSON(fiber.Map{"message": "AudioTrimmer executed"})
}
