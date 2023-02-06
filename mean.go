package finlib

import (
	"math"
	"math/big"
)

// Returns the mean of slice s.
func Mean(s []float64, p Precision) float64 {
	if len(s) == 0 {
		return math.NaN()
	}
	switch p {
	case Default:
		return kahanSum(s) / float64(len(s))
	case Naive:
		return nSum(s) / float64(len(s))
	case Exact:
		sum := bigSum(s)
		length := big.NewRat(int64(len(s)), 1)
		t := big.NewRat(0, 0).SetFloat64(sum)
		t = t.Quo(t, length)
		res, _ := t.Float64()
		return res
	case Fast:
		return nSum(s) / float64(len(s))
	default:
		sum := kahanSum(s)
		return sum / float64(len(s))
	}
}
