package emicalculator

import (
    "github.com/gofiber/fiber/v2"
    "strconv"
)

// @Summary Calculate EMI (Equated Monthly Installment)
// @Description Takes principal, rate, and tenure to calculate monthly EMI
// @Tags Tools
// @Accept json
// @Produce json
// @Param principal query number true "Principal amount"
// @Param rate query number true "Annual interest rate in percent"
// @Param tenure query integer true "Loan tenure in months"
// @Success 200 {object} map[string]float64
// @Failure 400 {object} map[string]string
// @Router /tools/emicalculator [get]
func Handler(c *fiber.Ctx) error {
    pStr := c.Query("principal")
    rateStr := c.Query("rate")
    tenureStr := c.Query("tenure")

    P, err1 := strconv.ParseFloat(pStr, 64)
    r, err2 := strconv.ParseFloat(rateStr, 64)
    n, err3 := strconv.Atoi(tenureStr)

    if err1 != nil || err2 != nil || err3 != nil || P < 0 || r < 0 || n <= 0 {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid principal, rate or tenure"})
    }

    emi := CalculateEMI(P, r, n)
    return c.JSON(fiber.Map{"emi": emi})
}
