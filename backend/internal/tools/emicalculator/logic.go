package emicalculator

import "math"

// CalculateEMI returns the Equated Monthly Instalment 
// P = principal, r = annual interest rate in percent, n = tenure in months
func CalculateEMI(P float64, annualRate float64, n int) float64 {
    r := annualRate / 1200 // monthly interest rate
    numerator := P * r * math.Pow(1+r, float64(n))
    denominator := math.Pow(1+r, float64(n)) - 1
    if denominator == 0 {
        return 0
    }
    return numerator / denominator
}
