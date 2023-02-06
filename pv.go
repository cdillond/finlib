package finlib

import "math"

// Returns the present value given a risk-free rate, the nubmer of periods, the fixed payment per period, and the future value.
// eop should be true when payments are made at the end of each period and false if payments are made at the beginning of each period.
func Pv(rate, nper, pmt, fv float64, eop bool) float64 {
	if rate == 0 {
		return -1 * (pmt*nper + fv)
	}
	var tp float64
	if !eop {
		tp = 1
	}
	p1 := math.Pow(1+rate, nper)
	p2 := pmt * (1 + rate*tp)
	p3 := (p1 - 1) / rate
	return (fv + p2*p3) / (-1 * p1)
}
