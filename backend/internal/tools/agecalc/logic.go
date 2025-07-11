package agecalc

import "time"

func CalculateAge(birthYear int) int {
	currentYear := time.Now().Year()
	return currentYear - birthYear
}
