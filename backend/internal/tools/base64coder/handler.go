package base64coder

import (
    "github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx) error {
    mode := c.Query("mode", "encode")
    data := c.Query("data", "")

    switch mode {
    case "decode":
        decoded, err := Decode(data)
        if err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid base64 data"})
        }
        return c.JSON(fiber.Map{"decoded": decoded})
    default: // encode
        encoded := Encode(data)
        return c.JSON(fiber.Map{"encoded": encoded})
    }
}
