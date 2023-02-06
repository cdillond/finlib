package finlib

import (
	"math"
	"math/big"
)

// Returns the population variance of s.
func VarP(s []float64, p Precision) float64 {
	if len(s) == 0 {
		return math.NaN()
	}
	if len(s) == 1 {
		return 0
	}
	switch p {
	case Default:
		return ecVarP(s)
	case Naive:
		return nVarP(s)
	case Exact:
		return bigVarP(s)
	case Fast:
		return fastVarP(s)
	default:
		// should not happen
		return ecVarP(s)
	}
}

// Returns the sample variance of s.
func VarS(s []float64, p Precision) float64 {
	if len(s) < 2 {
		return math.NaN()
	}
	switch p {
	case Default:
		return ecVarS(s)
	case Naive:
		return nVarS(s)
	case Exact:
		return bigVarS(s)
	case Fast:
		return fastVarS(s)
	default:
		// should not happen
		return ecVarS(s)
	}
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

// two-pass algorithm; should be exact but slow
func bigVarP(s []float64) float64 {
	b := make([]*big.Rat, len(s))
	sum := big.NewRat(0, 1)
	for i := 0; i < len(b); i++ {
		f := big.NewRat(0, 1).SetFloat64(s[i])
		b[i] = f
		sum = sum.Add(sum, f)
	}
	mean := sum.Quo(sum, big.NewRat(int64(len(s)), 1))
	difs := big.NewRat(0, 1)
	for i := 0; i < len(b); i++ {
		temp := big.NewRat(0, 1)
		temp = temp.Sub(mean, b[i])
		difs = difs.Add(difs, temp.Mul(temp, temp))
	}
	difs = difs.Quo(difs, big.NewRat(int64(len(s)), 1))
	out, _ := difs.Float64()
	return out
}

// two-pass algorithm; should be exact but slow
func bigVarS(s []float64) float64 {
	b := make([]*big.Rat, len(s))
	sum := big.NewRat(0, 1)
	for i := 0; i < len(b); i++ {
		f := big.NewRat(0, 1).SetFloat64(s[i])
		b[i] = f
		sum = sum.Add(sum, f)
	}
	mean := sum.Quo(sum, big.NewRat(int64(len(s)), 1))
	difs := big.NewRat(0, 1)
	for i := 0; i < len(b); i++ {
		temp := big.NewRat(0, 1)
		temp = temp.Sub(mean, b[i])
		difs = difs.Add(difs, temp.Mul(temp, temp))
	}
	difs = difs.Quo(difs, big.NewRat(int64(len(s))-1, 1))
	out, _ := difs.Float64()
	return out
}

// naive two-pass algorithm
func nVarP(s []float64) float64 {
	mean := Mean(s, Naive)
	var difs float64
	for i := 0; i < len(s); i++ {
		difs += (mean - s[i]) * (mean - s[i])
	}
	return difs / float64(len(s))
}

// naive two-pass algorithm
func nVarS(s []float64) float64 {
	mean := Mean(s, Naive)
	var difs float64
	for i := 0; i < len(s); i++ {
		difs += (mean - s[i]) * (mean - s[i])
	}
	return difs / (float64(len(s) - 1))
}

// naive one-pass algorithm
func fastVarP(s []float64) float64 {
	var sum, squaresum float64
	for i := 0; i < len(s); i++ {
		sum += s[i]
		squaresum += s[i] * s[i]
	}
	mean := sum / float64(len(s))
	return squaresum/float64(len(s)) - mean*mean
}

// naive one-pass sample variance algorithm
func fastVarS(s []float64) float64 {
	return fastVarP(s) * float64(len(s)) / float64(len(s)-1)
}
