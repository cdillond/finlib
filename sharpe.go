package finlib

import "math"

// Returns the Sharpe ratio of s, given the risk free rate.
func Sharpe(s []float64, rfr float64, p Precision) float64 {
	if len(s) < 1 {
		return math.NaN()
	}
	if s[0] == 0 {
		return math.NaN()
	}
	nom := s[len(s)-1]/s[0] - 1 - rfr
	std := StdP(s, p)
	if std == 0 {
		return math.NaN()
	}
	return nom / std
}
