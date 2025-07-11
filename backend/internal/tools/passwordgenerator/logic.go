package passwordgenerator

import (
    "math/rand"
    "time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
    "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
    rand.Seed(time.Now().UnixNano())
}

func Generate(length int) string {
    b := make([]byte, length)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
    }
    return string(b)
}
