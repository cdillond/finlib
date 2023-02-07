package finlib

// Returns the beta of s1 with respect to the benchmark s2.
func Beta(s1, s2 []float64) float64 {
	// this could also be attained using the CovMP func, but this is possibly more numerically stable
	v, c := VarP(s1), CoVarP(s1, s2)
	return c / v
}
