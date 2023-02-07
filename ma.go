package finlib

import (
	"math"
)

// Returns the simple moving averages of each period-length window of s.
// To avoid O(n^2) time complexity, Sma is calculated on a rolling basis.
// As a result of this, if s includes a NaN or Inf value anywhere, at least part of the return value will be corrupted.
func Sma(s []float64, period int) []float64 {
	res := make([]float64, len(s))
	// return early if period is 0 or out of bounds
	if period < 1 || period > len(s) {
		for i := 0; i < len(res); i++ {
			res[i] = math.NaN()
		}
		return res
	}
	p := float64(period) // typecast only once
	var sum float64
	// set all values below res[period-1] to NaN; accumulate sum
	for i := 0; i < period-1; i++ {
		res[i] = math.NaN()
		sum += s[i]
	}
	// set res[period-1] to sum / period;
	// this first one needs to happen separately to avoid having to bounds-check on each access to s[i-period]
	sum += s[period-1]
	res[period-1] = sum / p
	// proceed normally
	for i := period; i < len(s); i++ {
		sum += s[i] - s[i-period] // add leading edge, subtract trailing edge
		res[i] = sum / p
	}
	return res
}

// Returns the weighted moving average of s; implemented in O(n^2) time complexity.
func Wma(s, weights []float64) []float64 {
	res := make([]float64, len(s))
	if len(weights) >= len(s) || len(weights) < 1 {
		return res
	}
	if len(s) == 1 {
		res[0] = s[0]
		return res
	}
	p := float64(len(weights))
	for i := 0; i < len(weights)-1; i++ {
		res[i] = math.NaN()
	}

	for i := len(weights) - 1; i < len(s); i++ {
		k, _ := Mul(weights, s[i-len(weights):i])
		res[i] = Sum(k) / p
	}
	return res
}

// Returns the exponential moving average of s.
func Ema(s []float64, smoothing int) []float64 {
	res := make([]float64, len(s))
	if smoothing < 1 || smoothing > len(s) {
		for i := 0; i < len(res); i++ {
			res[i] = math.NaN()
		}
		return res
	}

	// calculate Sma for initial val (res[smoothing-1])
	var sum float64
	for i := 0; i < smoothing-1; i++ {
		sum += s[i]
		res[i] = math.NaN()
	}
	k := float64(smoothing)
	K := 2 / (k + 1)
	res[smoothing-1] = (sum + s[smoothing-1]) / k
	prevEma := res[smoothing-1]

	for i := smoothing; i < len(s); i++ {
		e := K*(s[i]-prevEma) + prevEma
		res[i] = e
		prevEma = e
	}
	return res
}
