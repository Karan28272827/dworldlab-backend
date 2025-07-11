package imagecropper

import "github.com/gofiber/fiber/v2"

func Handler(c *fiber.Ctx) error {
	// TODO: implement ImageCropper logic
	return c.JSON(fiber.Map{"message": "ImageCropper executed"})
}
