package finlib

import "math"

// Returns the effective annual interest rate, given the nominal annual interest rate (nom) and the number of compounding periods per year (f).
func Effect(nom, f float64) float64 {
	return math.Pow(1+(nom/f), f) - 1
}
