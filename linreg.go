package finlib

import (
	"math"
)

// Returns a simple linear regression model for s using the ordinary least squares method.
// The returned slice corresponds to the model y = mx + b, where return[0] = b and return[1] = m.
func LinReg(s []float64, p Precision) []float64 {
	switch p {
	case Default:
		return ecLinReg(s)
	case Exact:
		return ecLinReg(s)

	}
	var sumx, sumy, sumxy, sumxx float64
	var j float64
	for i := 0; i < len(s); i, j = i+1, j+1 {
		sumx += j
		sumy += s[i]
		sumxy += s[i] * j
		sumxx += j * j
	}
	fl := float64(len(s))
	dena := (fl*sumxx - sumx*sumx)
	if dena == 0 || len(s) == 0 {
		return []float64{math.NaN(), math.NaN()}
	}
	a := (fl*sumxy - sumx*sumy) / dena
	b := sumy/fl - a*sumx/fl
	return []float64{b, a}
}

func ecLinReg(s []float64) []float64 {
	var sumx, sumy, sumxy, sumxx float64
	var j float64
	var sumxi, sumxxi int
	var c1, c2, y1, y2, t1, t2 float64
	n := len(s) //- 1
	// https://en.wikipedia.org/wiki/Arithmetic_progression
	// first term = 0; last term = n - 1; there are n terms
	sumxi = n * (n - 1) / 2

	// ∑i^2(from i=1 to n) = n(n+1)(2n+1)/6 // this is inclusive of n; therefore, we can correct it by subtracting n^2
	sumxxi = n*(n+1)*(2*n+1)/6 - n*n
	for i := 0; i < len(s); i, j = i+1, j+1 {
		y1 = s[i] - c1
		t1 = sumy + y1
		c1 = t1 - sumy - y1
		sumy = t1

		y2 = s[i]*j - c2
		t2 = sumxy + y2
		c2 = t2 - sumxy - y2
		sumxy = t2
	}

	sumx = float64(sumxi)
	sumxx = float64(sumxxi)

	fl := float64(len(s))
	dena := (fl*sumxx - sumx*sumx)
	if dena == 0 || len(s) == 0 {
		return []float64{math.NaN(), math.NaN()}
	}
	a := (fl*sumxy - sumx*sumy) / dena
	b := sumy/fl - a*sumx/fl
	return []float64{b, a}
}
