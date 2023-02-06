package finlib

import (
	"math"
	"sort"
)

// Returns the median value of s. If s contains NaNs, they are treated as the lowest values of s.
func Median(s []float64) float64 {
	if len(s) == 0 {
		return math.NaN()
	}
	if len(s) == 1 {
		return s[0]
	}

	c := make([]float64, len(s))
	// necessary to avoid rearranging s when the slice is sorted
	copy(c, s)

	sort.Float64s(c)
	if len(c)%2 == 0 {
		return (c[(len(s)/2)-1] + c[len(s)/2]) / 2
	}
	return c[len(s)/2]
}
