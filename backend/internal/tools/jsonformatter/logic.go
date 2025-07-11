package jsonformatter

import (
	"bytes"
	"encoding/json"
)

// FormatJSON indents or minifies a JSON string.
// mode="pretty" → indented, mode="compact" → minified.
func FormatJSON(input string, mode string) (string, error) {
	var buf bytes.Buffer
	if mode == "compact" {
		if err := json.Compact(&buf, []byte(input)); err != nil {
			return "", err
		}
		return buf.String(), nil
	}
	// default to pretty
	if err := json.Indent(&buf, []byte(input), "", "  "); err != nil {
		return "", err
	}
	return buf.String(), nil
}
