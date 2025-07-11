package imageconverter

import "github.com/gofiber/fiber/v2"

func Handler(c *fiber.Ctx) error {
	// TODO: implement ImageConverter logic
	return c.JSON(fiber.Map{"message": "ImageConverter executed"})
}
