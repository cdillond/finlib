package finlib

// Returns the mean of slice s.
func Mean(s []float64) float64 {
	return kahanSum(s) / float64(len(s))
}
