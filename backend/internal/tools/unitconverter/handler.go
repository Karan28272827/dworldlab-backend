package unitconverter

import (
    "github.com/gofiber/fiber/v2"
    "strconv"
)

// @Summary Convert between different units
// @Description Convert values between length, weight, and temperature units
// @Tags Tools
// @Accept json
// @Produce json
// @Param value query number true "Value to convert"
// @Param from query string true "Source unit (km, m, kg, g, c, f)"
// @Param to query string true "Target unit (km, m, kg, g, c, f)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /tools/unitconverter [get]
func Handler(c *fiber.Ctx) error {
    valStr := c.Query("value")
    from := c.Query("from")
    to := c.Query("to")

    v, err1 := strconv.ParseFloat(valStr, 64)
    if err1 != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid value"})
    }

    result, err2 := Convert(v, from, to)
    if err2 != nil {
        return c.Status(400).JSON(fiber.Map{"error": err2.Error()})
    }

    return c.JSON(fiber.Map{
        "from":   from,
        "to":     to,
        "input":  v,
        "output": result,
    })
}
