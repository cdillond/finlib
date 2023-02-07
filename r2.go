package finlib

import "math"

// Returns the coefficient of determination, or R squared, of s1 with respect to s2.
func R2(s1, s2 []float64) float64 {
	// naive 2 pass algorithm
	mean := Mean(s1)
	var ssr, sst float64
	// not sure if this is numerically stable
	for i := 0; i < len(s1); i++ {
		ssr += (s1[i] - s2[i]) * (s1[i] - s2[i])
		sst += (s1[i] - mean) * (s1[i] - mean)
	}
	if sst == 0 {
		return math.NaN()
	}
	return 1 - (ssr / sst)
}
