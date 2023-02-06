package finlib

// Returns the moving average convergence-divergence of s. Conventionally, fast = 9, slow = 26, and signal = 12.
// The return values are, in order: fast EMA, slow EMA, MACD, and the signal line.
func Macd(s []float64, fast, slow, signal int, p Precision) ([]float64, []float64, []float64, []float64) {
	fa := Ema(s, fast)
	sl := Ema(s, slow)
	macd := Sub(fa, sl)
	si := Ema(macd, signal)
	return fa, sl, macd, si
}
