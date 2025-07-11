package colorpicker

import (
    "github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx) error {
    hex := c.Query("hex", "#000000")
    mode := c.Query("mode", "toRGB")

    switch mode {
    case "toHex":
        // Expect r,g,b as query params
        r := c.QueryInt("r")
        g := c.QueryInt("g")
        b := c.QueryInt("b")
        hex := RGBToHex(RGB{r, g, b})
        return c.JSON(fiber.Map{"hex": hex})
    default: // toRGB
        rgb, err := HexToRGB(hex)
        if err != nil {
            return c.Status(400).JSON(fiber.Map{"error": err.Error()})
        }
        return c.JSON(fiber.Map{"hex": hex, "rgb": rgb})
    }
}
