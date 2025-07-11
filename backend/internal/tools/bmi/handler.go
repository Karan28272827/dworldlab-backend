package bmi

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Summary Calculate BMI
// @Description Takes weight (kg) and height (m), returns BMI value
// @Tags Tools
// @Accept json
// @Produce json
// @Param weight query number true "Weight in kg"
// @Param height query number true "Height in meters"
// @Success 200 {object} map[string]float64
// @Failure 400 {object} map[string]string
// @Router /tools/bmi [get]

func Handler(c *fiber.Ctx) error {
	weightStr := c.Query("weight")
	heightStr := c.Query("height")

	weight, err1 := strconv.ParseFloat(weightStr, 64)
	height, err2 := strconv.ParseFloat(heightStr, 64)

	if err1 != nil || err2 != nil || height <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid weight or height"})
	}

	bmi := weight / (height * height)

	return c.JSON(fiber.Map{
		"bmi": bmi,
	})
}
