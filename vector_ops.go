package finlib

import (
	"errors"
	"math"
)

var ErrDifLen = errors.New("slices s1 and s2 must be of equal length")

// Returns the element-wise sum of s1 and s2. s1 and s2 must have equal lengths.
func Add(s1, s2 []float64) ([]float64, error) {
	l1, l2 := len(s1), len(s2)
	if l1 != l2 {
		return []float64{}, ErrDifLen
	}
	res := make([]float64, l1)
	for i := 0; i < l1; i++ {
		res[i] = s1[i] + s2[i]
	}
	return res, nil
}

// Returns the element-wise quotient of s1 divided by s2. s1 and s2 must have equal lengths.
func Div(s1, s2 []float64) ([]float64, error) {
	l1, l2 := len(s1), len(s2)
	if l1 != l2 {
		return []float64{}, ErrDifLen
	}
	res := make([]float64, l1)
	for i := 0; i < l1; i++ {
		if s2[i] == 0 {
			res[i] = math.NaN()
		}
		res[i] = s1[i] / s2[i]
	}
	return res, nil
}

// Returns the element-wise product (Hadamard product) of s1 and s2. s1 and s2 must have equal lengths.
func Mul(s1, s2 []float64) ([]float64, error) {
	l1, l2 := len(s1), len(s2)
	if l1 != l2 {
		return []float64{}, ErrDifLen
	}
	res := make([]float64, l1)
	for i := 0; i < l1; i++ {
		res[i] = s1[i] * s2[i]
	}
	return res, nil
}

// Returns the element-wise difference of s2 subtracted from s1. s1 and s2 must have equal lengths.
func Sub(s1, s2 []float64) ([]float64, error) {
	l1, l2 := len(s1), len(s2)
	if l1 != l2 {
		return []float64{}, ErrDifLen
	}
	res := make([]float64, l1)
	for i := 0; i < l1; i++ {
		res[i] = s1[i] - s2[i]
	}
	return res, nil
}
