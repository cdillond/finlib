package finlib

import (
	"math"
)

// Returns s1[i] + s2[i]
func Add(s1, s2 []float64) []float64 {
	l1, l2 := len(s1), len(s2)
	if l1 != l2 {
		return []float64{}
	}
	res := make([]float64, l1)
	for i := 0; i < l1; i++ {
		res[i] = s1[i] + s2[i]
	}
	return res
}

// Returns s1[i]/s2[i]
func Div(s1, s2 []float64) []float64 {
	l1, l2 := len(s1), len(s2)
	if l1 != l2 {
		return []float64{}
	}
	res := make([]float64, l1)
	for i := 0; i < l1; i++ {
		if s2[i] == 0 {
			res[i] = math.NaN()
		}
		res[i] = s1[i] / s2[i]
	}
	return res
}

// Returns s1[i]*s2[i]
func Mul(s1, s2 []float64) []float64 {
	l1, l2 := len(s1), len(s2)
	if l1 != l2 {
		return []float64{}
	}
	res := make([]float64, l1)
	for i := 0; i < l1; i++ {
		res[i] = s1[i] * s2[i]
	}
	return res
}

// Returns s1-s2
func Sub(s1, s2 []float64) []float64 {
	l1, l2 := len(s1), len(s2)
	if l1 != l2 {
		return []float64{}
	}
	res := make([]float64, l1)
	for i := 0; i < l1; i++ {
		res[i] = s1[i] - s2[i]
	}
	return res
}
