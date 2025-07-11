package wordcounter

import "strings"

func Count(text string) int {
    fields := strings.Fields(text)
    return len(fields)
}
