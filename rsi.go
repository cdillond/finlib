package finlib

// Returns the relative strength index of s, given the specified period.
func Rsi(s []float64, period int) []float64 {
	ups, downs := make([]float64, len(s)), make([]float64, len(s))
	for i := 1; i < len(s); i++ {
		delta := s[i] - s[i-1]
		if delta < 0 {
			delta *= -1
			downs[i] = delta
		} else {
			ups[i] = delta
		}
	}
	uEma := Ema(ups, period)
	dEma := Ema(downs, period)
	rsi := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		//if dEma[i] == 0 {
		//	rsi[i] = math.NaN()
		//} else {
		rsi[i] = uEma[i] / (uEma[i] + dEma[i])
		//}
	}
	return rsi
}

// Returns the Cutler's RSI of s, given the specified period.
func CutlerRsi(s []float64, period int) []float64 {
	ups, downs := make([]float64, len(s)), make([]float64, len(s))
	for i := 1; i < len(s); i++ {
		delta := s[i] - s[i-1]
		if delta < 0 {
			delta *= -1
			downs[i] = delta
		} else {
			ups[i] = delta
		}
	}
	uSma := Sma(ups, period)
	dSma := Sma(downs, period)
	rsi := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		//if dSma[i] == 0 {
		//	rsi[i] = math.NaN()
		//} else {
		rsi[i] = uSma[i] / (uSma[i] + dSma[i])
		//}
	}
	return rsi
}
