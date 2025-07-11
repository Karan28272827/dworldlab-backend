package texttospeech

import "github.com/gofiber/fiber/v2"

func Handler(c *fiber.Ctx) error {
	// TODO: implement TextToSpeech logic
	return c.JSON(fiber.Map{"message": "TextToSpeech executed"})
}
