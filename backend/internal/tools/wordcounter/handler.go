package wordcounter

import (
    "github.com/gofiber/fiber/v2"
)

// @Summary Count words in text
// @Description Count the number of words in the provided text
// @Tags Tools
// @Accept json
// @Produce json
// @Param text query string false "Text to count words in"
// @Success 200 {object} map[string]int
// @Router /tools/wordcounter [get]
func Handler(c *fiber.Ctx) error {
    text := c.Query("text", "")
    count := Count(text)
    return c.JSON(fiber.Map{"count": count})
}
