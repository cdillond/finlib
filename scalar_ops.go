package finlib

import (
	"math"
)

// Adds f to each value in s.
func SAdd(f float64, s []float64) []float64 {
	res := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = s[i] + f
	}
	return res
}

// Divides each value of s by f.
func SDiv(f float64, s []float64) []float64 {
	res := make([]float64, len(s))
	if f == 0 {
		for i := 0; i < len(s); i++ {
			res[i] = math.NaN()
		}
	}
	for i := 0; i < len(s); i++ {
		res[i] = s[i] / f
	}
	return res
}

// Multiplies f by each value of s.
func SMul(f float64, s []float64) []float64 {
	res := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = s[i] * f
	}
	return res
}

// Subtracts f from each value of s.
func SSub(f float64, s []float64) []float64 {
	res := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = s[i] - f
	}
	return res
}
