package finlib

import "math"

// Returns the Pearson correlation coefficient for the populations s1 and s2.
func CorP(s1, s2 []float64) float64 {
	cvm := CovMP(s1, s2)
	den := math.Sqrt(cvm[0][0]) * math.Sqrt(cvm[1][1])
	if den == 0 {
		return math.NaN()
	}
	return cvm[0][1] / den
}

// Returns the Pearson correlation coefficient for the samples s1 and s2.
func CorS(s1, s2 []float64) float64 {
	cvm := CovMS(s1, s2)
	den := math.Sqrt(cvm[0][0]) * math.Sqrt(cvm[1][1])
	if den == 0 {
		return math.NaN()
	}
	return cvm[0][1] / den
}
