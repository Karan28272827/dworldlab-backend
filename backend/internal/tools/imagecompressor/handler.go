package imagecompressor

import "github.com/gofiber/fiber/v2"

func Handler(c *fiber.Ctx) error {
	// TODO: implement ImageCompressor logic
	return c.JSON(fiber.Map{"message": "ImageCompressor executed"})
}
