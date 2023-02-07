package finlib

import (
	"math"
)

// Returns the population covariance of s1 and s2.
func CoVarP(s1, s2 []float64) float64 {
	if len(s1) != len(s2) || len(s1) < 2 {
		return math.NaN()
	}
	return ecCoVarP(s1, s2)
}

// this does multiple passes... https://www.osti.gov/servlets/purl/1028931
func ecCoVarP(s1, s2 []float64) float64 {
	var y1, c1, t1, sum1 float64
	var y2, c2, t2, sum2 float64

	// NOT SHIFTED
	for i := 0; i < len(s1); i++ {
		y1 = s1[i] - c1
		t1 = sum1 + y1
		c1 = t1 - sum1 - y1
		sum1 = t1

		y2 = s2[i] - c2
		t2 = sum2 + y2
		c2 = t2 - sum2 - y2
		sum2 = t2
	}
	m1 := sum1 / float64(len(s1))
	m2 := sum2 / float64(len(s2))

	var Y, C, T, Sum float64
	for i := 0; i < len(s1); i++ {
		Y = (s1[i]-m1)*(s2[i]-m2) - C
		T = Sum + Y
		C = T - Sum - Y
		Sum = T
	}
	return Sum / float64(len(s1))
}

// Returns the sample covariance of s1 and s2.
func CoVarS(s1, s2 []float64) float64 {
	if len(s1) != len(s2) || len(s1) < 2 {
		return math.NaN()
	}
	return ecCovS(s1, s2)
}

// this does multiple passes...
func ecCovS(s1, s2 []float64) float64 {
	var y1, c1, t1, sum1 float64
	var y2, c2, t2, sum2 float64

	// NOT SHIFTED
	for i := 0; i < len(s1); i++ {
		y1 = s1[i] - c1
		t1 = sum1 + y1
		c1 = t1 - sum1 - y1
		sum1 = t1

		y2 = s2[i] - c2
		t2 = sum2 + y2
		c2 = t2 - sum2 - y2
		sum2 = t2
	}
	m1 := sum1 / float64(len(s1))
	m2 := sum2 / float64(len(s2))

	var Y, C, T, Sum float64
	for i := 0; i < len(s1); i++ {
		Y = (s1[i]-m1)*(s2[i]-m2) - C
		T = Sum + Y
		C = T - Sum - Y
		Sum = T
	}
	return Sum / float64(len(s1)-1)
}
