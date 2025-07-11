package colorpicker

import (
    "errors"
    "fmt"
    "strconv"
    "strings"
)

// RGB holds red, green, blue values.
type RGB struct {
    R, G, B int
}

// HexToRGB converts a hex color (e.g. "#1a2b3c" or "1a2b3c") to RGB.
func HexToRGB(hex string) (RGB, error) {
    h := strings.TrimPrefix(hex, "#")
    if len(h) != 6 {
        return RGB{}, errors.New("invalid hex length")
    }
    r, err1 := strconv.ParseInt(h[0:2], 16, 0)
    g, err2 := strconv.ParseInt(h[2:4], 16, 0)
    b, err3 := strconv.ParseInt(h[4:6], 16, 0)
    if err1 != nil || err2 != nil || err3 != nil {
        return RGB{}, errors.New("invalid hex value")
    }
    return RGB{int(r), int(g), int(b)}, nil
}

// RGBToHex converts RGB values to a hex string.
func RGBToHex(c RGB) string {
    return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}
