package sipcalculator

import "math"

// CalculateSIP returns corpus accumulated for monthly investment P,
// at annualRate (%) for n months, compounded monthly.
func CalculateSIP(P float64, annualRate float64, n int) float64 {
    r := annualRate / 1200 // monthly rate
    factor := math.Pow(1+r, float64(n)) - 1
    if r == 0 {
        return P * float64(n)
    }
    return P * factor * (1 + r) / r
}
