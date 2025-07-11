package passwordgenerator

import (
    "github.com/gofiber/fiber/v2"
    "strconv"
)

// @Summary Generate random password
// @Description Generate a random password of specified length
// @Tags Tools
// @Accept json
// @Produce json
// @Param length query integer false "Password length (default: 12)"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /tools/passwordgenerator [get]
func Handler(c *fiber.Ctx) error {
    lengthStr := c.Query("length", "12")
    length, err := strconv.Atoi(lengthStr)
    if err != nil || length <= 0 {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid length"})
    }

    pwd := Generate(length)
    return c.JSON(fiber.Map{"password": pwd})
}
