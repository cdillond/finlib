package finlib

import (
	"math"
)

// Returns the beta of s1 with respect to the benchmark s2.
func Beta(s1, s2 []float64, p Precision) float64 {
	switch p {
	case Default:
		return ecBeta(s1, s2, p)
	case Naive:
		cvm := CovMP(s1, s2, p)
		if cvm[0][0] == 0 {
			return math.NaN()
		}
		return cvm[0][1] / cvm[0][0]
	case Exact:
		return ecBeta(s1, s2, p)
	case Fast:
		cvm := CovMP(s1, s2, p)
		if cvm[0][0] == 0 {
			return math.NaN()
		}
		return cvm[0][1] / cvm[0][0]
	default:
		return ecBeta(s1, s2, p)

	}
}

// Returns the beta of s1 with respect to the benchmark s2.
func ecBeta(s1, s2 []float64, p Precision) float64 {
	v, c := VarP(s1, p), CoVarP(s1, s2, p)
	if v == 0 {
		return math.NaN()
	}
	return c / v
}
