package speechtotext

import "github.com/gofiber/fiber/v2"

func Handler(c *fiber.Ctx) error {
	// TODO: implement SpeechToText logic
	return c.JSON(fiber.Map{"message": "SpeechToText executed"})
}
