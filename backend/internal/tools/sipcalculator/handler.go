package sipcalculator

import (
    "github.com/gofiber/fiber/v2"
    "strconv"
)

// @Summary Calculate SIP (Systematic Investment Plan) corpus
// @Description Takes monthly amount, rate, and months to calculate accumulated corpus
// @Tags Tools
// @Accept json
// @Produce json
// @Param amount query number true "Monthly investment amount"
// @Param rate query number true "Annual interest rate in percent"
// @Param months query integer true "Investment period in months"
// @Success 200 {object} map[string]float64
// @Failure 400 {object} map[string]string
// @Router /tools/sipcalculator [get]
func Handler(c *fiber.Ctx) error {
    pStr := c.Query("amount")
    rateStr := c.Query("rate")
    monthsStr := c.Query("months")

    P, err1 := strconv.ParseFloat(pStr, 64)
    r, err2 := strconv.ParseFloat(rateStr, 64)
    n, err3 := strconv.Atoi(monthsStr)

    if err1 != nil || err2 != nil || err3 != nil || P < 0 || r < 0 || n <= 0 {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid amount, rate or months"})
    }

    corpus := CalculateSIP(P, r, n)
    return c.JSON(fiber.Map{"corpus": corpus})
}
