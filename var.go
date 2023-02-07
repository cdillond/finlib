package finlib

import (
	"math"
)

// Returns the population variance of s.
func VarP(s []float64) float64 {
	if len(s) == 0 {
		return math.NaN()
	}
	return ecVarP(s)
}

// Returns the sample variance of s.
func VarS(s []float64) float64 {
	if len(s) < 2 {
		return math.NaN()
	}
	return ecVarS(s)
}

// error-corrected one-pass population variance algorithm
func ecVarP(s []float64) float64 {
	var y, c, t, sum float64        // used in summation part
	var Y, C, T, SquaresSum float64 // used in multipliction part
	K := s[0]                       // shift
	for i := 0; i < len(s); i++ {
		y = s[i] - K - c  // adjust s[i] based on shift and prev error
		t = sum + y       // add corrected s[i] to the sum to get a new sum
		c = (t - sum) - y // subtract old sum and y from the new sum to get the error
		sum = t           // assign new sum to old sum variable

		Y = (s[i]-K)*(s[i]-K) - C // Y = shifted mult - error
		T = SquaresSum + Y        // add corrected mult to the Sum to get a new Sum
		C = T - SquaresSum - Y    // subtract old Sum from the new Sum to get the error
		SquaresSum = T            // assign new Sum to old Sum variable
	}
	shiftedMean := (sum / float64(len(s))) // sum / len(s)
	// pop var = (1/N)*(Sum from i to N of xi^2) - mean^2
	return SquaresSum/float64(len(s)) - shiftedMean*shiftedMean
}

// error-corrected one-pass sample variance algorithm
func ecVarS(s []float64) float64 {
	return ecVarP(s) * float64(len(s)) / float64(len(s)-1)
}
