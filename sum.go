package finlib

import (
	"math/big"
)

// Returns the sum of all values in s.
func Sum(s []float64, p Precision) float64 {
	switch p {
	case Default:
		return kahanSum(s)
	case Naive:
		return nSum(s)
	case Exact:
		return bigSum(s)
	case Fast:
		return nSum(s)
	default:
		return kahanSum(s)
	}
}

func nSum(s []float64) float64 {
	var sum float64
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return sum
}

// error correcting
func kahanSum(s []float64) float64 {
	var sum float64
	var c float64 // correction
	for i := 0; i < len(s); i++ {
		y := s[i] - c   // adjust s[i] based on previous error
		t := sum + y    // add corrected s[i] to the sum to get a new sum
		c = t - sum - y // subtract old sum and y from the new sum to get the error
		sum = t         // assign new sum to old sum variable
	}
	return sum
}

func bigSum(s []float64) float64 {
	Sum := big.NewRat(0, 1)
	for i := 0; i < len(s); i++ {
		Sum = Sum.Add(Sum, big.NewRat(0, 1).SetFloat64(s[i]))
	}
	fsum, _ := Sum.Float64()
	return fsum
}
