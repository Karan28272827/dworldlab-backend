package agecalc

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Summary Calculate age from birth year
// @Description Takes birth year and returns current age
// @Tags Tools
// @Accept json
// @Produce json
// @Param year query integer true "Birth year"
// @Success 200 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Router /tools/age [get]
func Handler(c *fiber.Ctx) error {
	yearStr := c.Query("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid year"})
	}

	age := CalculateAge(year)

	return c.JSON(fiber.Map{"age": age})
}
