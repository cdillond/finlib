package finlib

import "math"

// Returns (s[i]/s[i-1] - 1) for each integer i in (0, len(s)).
func Roc(s []float64) []float64 {
	res := make([]float64, len(s))
	if len(s) < 1 {
		return res
	}
	res[0] = 0
	for i := 1; i < len(s); i++ {
		if s[i-1] == 0 {
			res[i] = math.NaN()
			continue
		}
		res[i] = (s[i] - s[i-1]) / s[i-1]
	}
	return res
}

// Returns s[i]/s[i-1] for each integer i in (0, len(s)).
func RocR(s []float64) []float64 {
	res := make([]float64, len(s))
	if len(s) < 1 {
		return res
	}
	res[0] = 0
	for i := 1; i < len(s); i++ {
		if s[i-1] == 0 {
			res[i] = math.NaN()
			continue
		}
		res[i] = s[i] / s[i-1]
	}
	return res
}
