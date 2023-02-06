package finlib

import (
	"math"
	"math/big"
)

// https://www.osti.gov/servlets/purl/1028931
func CoVarP(s1, s2 []float64, p Precision) float64 {
	if len(s1) != len(s2) || len(s1) < 2 {
		return math.NaN()
	}
	switch p {
	case Default:
		return ecCoVarP(s1, s2)
	case Naive:
		return nCoVarP(s1, s2)
	case Exact:
		return ecCoVarP(s1, s2)
		// return bigCoVarP(s1, s2) DOES NOT WORK
	case Fast:
		return nCoVarP(s1, s2)
	default:
		return ecCoVarP(s1, s2)
	}
}

// this does multiple passes...
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

// this does multiple passes...
func nCoVarP(s1, s2 []float64) float64 {
	var sum1, sum2 float64

	// NOT SHIFTED
	for i := 0; i < len(s1); i++ {
		sum1 += s1[i]
		sum2 += s2[i]
	}
	m1 := sum1 / float64(len(s1))
	m2 := sum2 / float64(len(s2))

	var sum float64
	for i := 0; i < len(s1); i++ {
		sum += (s1[i] - m1) * (s2[i] - m2)
	}
	return sum / float64(len(s1))
}

func bigCoVarP(s1, s2 []float64) float64 {
	sum1 := big.NewRat(0, 1)
	sum2 := big.NewRat(0, 1)

	length := big.NewRat(int64(len(s1)), 1)

	// NOT SHIFTED
	for i := 0; i < len(s1); i++ {
		a1, a2 := big.NewRat(0, 1), big.NewRat(0, 1)
		a1.SetFloat64(s1[i])
		a2.SetFloat64(s2[i])
		sum1.Add(sum1, a1)
		sum1.Add(sum2, a2)
	}
	m1 := sum1.Quo(sum1, length)
	m2 := sum2.Quo(sum2, length)

	sum := big.NewRat(0, 1)
	a, b := big.NewRat(0, 1), big.NewRat(0, 1)
	for i := 0; i < len(s1); i++ {
		a.SetFloat64(s1[i])
		b.SetFloat64(s2[i])
		sum.Add(sum, a.Mul(a.Sub(a, m1), b.Sub(b, m2)))
	}
	sum.Quo(sum, length)
	res, _ := sum.Float64()

	return res
}

// https://www.osti.gov/servlets/purl/1028931
func CoVarS(s1, s2 []float64, p Precision) float64 {
	if len(s1) != len(s2) || len(s1) < 2 {
		return math.NaN()
	}

	switch p {
	case Default:
		return ecCovS(s1, s2)
	case Naive:
		return nCovS(s1, s2)
	case Exact:
		return ecCovS(s1, s2)
		// return bigCovS(s1, s2) MIGHT NOT WORK
	case Fast:
		return nCovS(s1, s2)
	default:
		return ecCovS(s1, s2)
	}
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

func nCovS(s1, s2 []float64) float64 {
	var sum1, sum2 float64

	// NOT SHIFTED
	for i := 0; i < len(s1); i++ {
		sum1 += s1[i]
		sum2 += s2[i]
	}
	m1 := sum1 / float64(len(s1))
	m2 := sum2 / float64(len(s2))

	var sum float64
	for i := 0; i < len(s1); i++ {
		sum += (s1[i] - m1) * (s2[i] - m2)
	}
	return sum / float64(len(s1)-1)
}

func bigCovS(s1, s2 []float64) float64 {
	sum1 := big.NewRat(0, 1)
	sum2 := big.NewRat(0, 1)

	// NOT SHIFTED
	for i := 0; i < len(s1); i++ {
		sum1 = sum1.Add(sum1, big.NewRat(0, 1).SetFloat64(s1[i]))
		sum2 = sum2.Add(sum2, big.NewRat(0, 1).SetFloat64(s2[i]))
	}
	m1 := sum1.Quo(sum1, big.NewRat(int64(len(s1)), 1))
	m2 := sum2.Quo(sum2, big.NewRat(int64(len(s2)), 1))

	sum := big.NewRat(0, 1)
	for i := 0; i < len(s1); i++ {
		a := big.NewRat(0, 1).SetFloat64(s1[i])
		b := big.NewRat(0, 1).SetFloat64(s2[i])

		a = a.Sub(a, m1)
		b = b.Sub(b, m2)
		sum = sum.Add(sum, a.Mul(a, b))
	}
	sum = sum.Quo(sum, big.NewRat(int64(len(s1)-1), 1))
	res, _ := sum.Float64()
	return res
}
