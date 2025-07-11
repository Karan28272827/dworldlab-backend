package base64coder

import "encoding/base64"

// Encode returns the base64 encoding of data.
func Encode(data string) string {
    return base64.StdEncoding.EncodeToString([]byte(data))
}

// Decode returns the decoded string or an error.
func Decode(encoded string) (string, error) {
    b, err := base64.StdEncoding.DecodeString(encoded)
    return string(b), err
}
