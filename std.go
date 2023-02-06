package finlib

import "math"

// Returns the population standard deviation for s.
func StdP(s []float64, p Precision) float64 {
	return math.Sqrt(VarP(s, p))
}

// Returns the sample standard deviation for s.
func StdS(s []float64, p Precision) float64 {
	return math.Sqrt(VarS(s, p))
}
