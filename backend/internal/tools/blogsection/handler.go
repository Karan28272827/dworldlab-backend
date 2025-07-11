package blogsection

import "github.com/gofiber/fiber/v2"

func Handler(c *fiber.Ctx) error {
	// TODO: implement BlogSection logic
	return c.JSON(fiber.Map{"message": "BlogSection executed"})
}
