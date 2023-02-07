package finlib

import (
	"math"
)

// Replaces all NaN and Inf values with 0.
func ZeroNaNInfs(s []float64) []float64 {
	res := make([]float64, len(s))
	for i, f := range s {
		if math.IsNaN(f) || math.IsInf(f, 0) {
			continue
		}
		res[i] = f
	}
	return res
}

// Returns a copy of s with all values of s that are equivalent to a replaced with b.
// Note: all NaNs are treated as equivalent.
func Replace(a, b float64, s []float64) []float64 {
	res := make([]float64, len(s))
	switch {
	case math.IsNaN(a):
		for i, f := range s {
			if math.IsNaN(f) {
				res[i] = b
				continue
			}
			res[i] = f
		}
		return res
	case math.IsInf(a, -1):
		for i, f := range s {
			if math.IsInf(f, -1) {
				res[i] = b
				continue
			}
			res[i] = f
		}
		return res
	case math.IsInf(a, 1):
		for i, f := range s {
			if math.IsInf(f, 1) {
				res[i] = b
				continue
			}
			res[i] = f
		}
		return res
	default:
		for i, f := range s {
			if f == a {
				res[i] = b
				continue
			}
			res[i] = f
		}
		return res
	}
}
