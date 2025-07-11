package timerstopwatch

import "github.com/gofiber/fiber/v2"

func Handler(c *fiber.Ctx) error {
	// TODO: implement TimerStopwatch logic
	return c.JSON(fiber.Map{"message": "TimerStopwatch executed"})
}
