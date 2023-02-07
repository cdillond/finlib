package finlib

// Returns a slice of return values for fn(s[i]) for each index (i) of s.
func MapTo(s []float64, fn func(float64) float64) []float64 {
	c := make([]float64, len(s))

	for i := 0; i < len(s); i++ {
		c[i] = fn(s[i])
	}
	return c
}

// Returns a slice of return values for fn(s[i], f) for each index (i) of s.
func MapFToS(s []float64, f float64, fn func(f1, f2 float64) float64) []float64 {
	c := make([]float64, len(s))

	for i := 0; i < len(s); i++ {
		c[i] = fn(f, s[i])
	}
	return c
}
