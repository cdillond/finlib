package finlib

import (
	"math"
)

// Returns the SMA with the specified period and the high and low Bollinger bands of the specified width.
// High = SMA+width*StdP(SMA), Low = SMA-width*StdP(SMA)
func Boll(s []float64, period int, width float64) ([]float64, []float64, []float64) {
	sma := Sma(s, period)
	hi, lo := make([]float64, len(s)), make([]float64, len(s))
	for i := 0; i < period; i++ {
		hi[i], lo[i] = math.NaN(), math.NaN()
	}
	for i := period; i < len(sma); i++ {
		hi[i] = sma[i] + width*StdP(s[i-period:i], 0)
		lo[i] = sma[i] - width*StdP(s[i-period:i], 0)
	}
	return sma, hi, lo
}
