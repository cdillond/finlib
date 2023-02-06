package finlib

import "math"

// Returns the effective annual interest rate, given the nominal annual interest rate and the number of compounding periods per year.
func Effect(nom, nPerY float64) float64 {
	return math.Pow(1+(nom/nPerY), nPerY) - 1
}
