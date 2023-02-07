package finlib

import (
	"math"
)

// Returns the population covariance matrix of s1 and s2. For the resulting slice r,
// r[0][0] = VarP(s1), r[0][1] = CovP(s1, s2), r[1][0] = CovP(s2, s1), r[1][1] = VarP(s2).
func CovMP(s1, s2 []float64) [][]float64 {
	if len(s1) != len(s2) || len(s1) < 1 {
		return [][]float64{{math.NaN(), math.NaN()}, {math.NaN(), math.NaN()}}
	}

	return ecCovMP(s1, s2)

}

// requires two passes
func ecCovMP(s1, s2 []float64) [][]float64 {
	var y1, c1, t1, sum1 float64
	var y2, c2, t2, sum2 float64

	var Y1, C1, T1, SquaresSum1 float64
	var Y2, C2, T2, SquaresSum2 float64

	K1, K2 := s1[0], s2[0]
	for i := 0; i < len(s1); i++ {
		y1 = s1[i] - c1 - K1
		t1 = sum1 + y1
		c1 = t1 - sum1 - y1
		sum1 = t1

		Y1 = (s1[i]-K1)*(s1[i]-K1) - C1 // Y = shifted mult - error
		T1 = SquaresSum1 + Y1           // add corrected mult to the Sum to get a new Sum
		C1 = T1 - SquaresSum1 - Y1      // subtract old Sum from the new Sum to get the error
		SquaresSum1 = T1                // assign new Sum to old Sum variable

		y2 = s2[i] - c2 - K2
		t2 = sum2 + y2
		c2 = t2 - sum2 - y2
		sum2 = t2

		Y2 = (s2[i]-K2)*(s2[i]-K2) - C2 // Y = shifted mult - error
		T2 = SquaresSum2 + Y2           // add corrected mult to the Sum to get a new Sum
		C2 = T2 - SquaresSum2 - Y2      // subtract old Sum from the new Sum to get the error
		SquaresSum2 = T2                // assign new Sum to old Sum variable

	}
	m1 := sum1/float64(len(s1)) + K1
	m2 := sum2/float64(len(s2)) + K2

	var Y, C, T, Sum float64
	for i := 0; i < len(s1); i++ {
		Y = (s1[i]-m1)*(s2[i]-m2) - C
		T = Sum + Y
		C = T - Sum - Y
		Sum = T
	}
	var1 := SquaresSum1/float64(len(s1)) - (m1-K1)*(m1-K1)
	var2 := SquaresSum2/float64(len(s2)) - (m2-K2)*(m2-K2)
	cv := Sum / float64(len(s1))
	return [][]float64{{var1, cv}, {cv, var2}}
}

// Returns the sample covariance matrix of s1 and s2. For the resulting slice r,
// r[0][0] = VarS(s1), r[0][1] = CovS(s1, s2), r[1][0] = CovS(s2, s1), r[1][1] = VarS(s2).
func CovMS(s1, s2 []float64) [][]float64 {
	if len(s1) != len(s2) || len(s1) < 2 {
		return [][]float64{{math.NaN(), math.NaN()}, {math.NaN(), math.NaN()}}
	}

	matrix := ecCovMP(s1, s2)
	matrix[0][0] = matrix[0][0] * float64(len(s1)) / float64(len(s1)-1)
	matrix[0][1] = matrix[0][1] * float64(len(s1)) / float64(len(s1)-1)
	matrix[1][0] = matrix[1][0] * float64(len(s1)) / float64(len(s1)-1)
	matrix[1][1] = matrix[1][1] * float64(len(s1)) / float64(len(s1)-1)
	return matrix

}
