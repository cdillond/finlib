package finlib

import "math"

// Returns the future value given the interest rate per priod (rate), the number of periods (nper), a fixed payment amount (pmt), and the present value (pv).
// eop should be set to true if payments are made at the end of the period, and false if they are made at the beginning of the period.
func Fv(rate, nper, pmt, pv float64, eop bool) float64 {
	if rate == 0 {
		return -1 * (pmt*nper + pv)
	}
	var tp float64
	if !eop {
		tp = 1
	}
	p1 := math.Pow(1+rate, nper)
	p2 := pmt * (1 + rate*tp)
	p3 := (p1 - 1) / rate
	return -1 * (pv*p1 + p2*p3)
}
