package qrcodegenerator

import (
    "encoding/base64"

    "github.com/skip2/go-qrcode"
)

// GenerateQRCodePNG encodes the given text as a PNG and returns it base64‑encoded.
func GenerateQRCodePNG(text string, size int) (string, error) {
    png, err := qrcode.Encode(text, qrcode.Medium, size)
    if err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(png), nil
}
