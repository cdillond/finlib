package finlib

import (
	"math"
)

// Returns the maximum value in s.
func Max(s []float64) float64 {
	if len(s) < 1 {
		return math.NaN()
	}
	max := s[0]
	for i := 0; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}

// Returns the first index of the maximum value included in s.
func MaxIndex(s []float64) int {
	var index int
	if len(s) < 1 {
		return index
	}
	max := s[0]
	for i := 0; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
			index = i
		}
	}
	return index
}

// Returns the minimum value in s.
func Min(s []float64) float64 {
	if len(s) < 1 {
		return math.NaN()
	}
	min := s[0]
	for i := 0; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
	}
	return min
}

// Returns the first index of the minimum value included in s.
func MinIndex(s []float64) int {
	var index int
	if len(s) < 1 {
		return index
	}
	min := s[0]
	for i := 0; i < len(s); i++ {
		if s[i] > min {
			min = s[i]
			index = i
		}
	}
	return index
}

// Returns the minimum and maximum values included in s.
func MinMax(s []float64) (float64, float64) {
	if len(s) < 1 {
		return math.NaN(), math.NaN()
	}
	min, max := s[0], s[0]
	for i := 0; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
		if s[i] > max {
			max = s[i]
		}
	}
	return min, max
}

func MinMaxIndex(s []float64) (int, int) {
	minIndex, maxIndex := 0, 0
	if len(s) < 1 {
		return minIndex, maxIndex
	}
	min, max := s[0], s[0]
	for i := 0; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
			minIndex = i
		}
		if s[i] > max {
			max = s[i]
			maxIndex = i
		}
	}
	return minIndex, maxIndex
}
