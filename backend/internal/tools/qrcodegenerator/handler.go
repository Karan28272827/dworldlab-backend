package qrcodegenerator

import (
    "github.com/gofiber/fiber/v2"
    "strconv"
)

func Handler(c *fiber.Ctx) error {
    text := c.Query("text", "")
    sizeStr := c.Query("size", "256")
    size, err := strconv.Atoi(sizeStr)
    if err != nil || size <= 0 {
        size = 256
    }

    imgBase64, err := GenerateQRCodePNG(text, size)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to generate QR code"})
    }

    return c.JSON(fiber.Map{
        "text":      text,
        "size":      size,
        "imageBase64": imgBase64,
    })
}
