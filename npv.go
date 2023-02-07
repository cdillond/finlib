package finlib

import (
	"math"
	"time"
)

// Calculates the net present value of a series of cashflows at the given rate.
// The cashflows are considered consecutively.
func Npv(rate float64, cashflows []float64) float64 {
	if len(cashflows) == 0 {
		return math.NaN()
	}
	var res float64
	for i := 1; i < len(cashflows); i++ {
		res += cashflows[i] / math.Pow(1+rate, float64(i+1))
	}
	return res + cashflows[0]
}

// Caculates the net present value of a series of cashflows at a given date. The period is assumed to be 365 days.
// cashflows and dates must be of equal length. Each cashflow is expected to occur at the date stored at the corresponding index of dates.
func Xnpv(rate float64, cashflows []float64, dates []time.Time) float64 {
	var res float64
	if len(cashflows) != len(dates) || len(cashflows) == 0 {
		return math.NaN()
	}
	start := dates[0]
	for i := 1; i < len(cashflows); i++ {
		t1, t2 := start.Truncate(time.Hour), dates[i].Truncate(time.Hour)
		if t1.After(t2) {
			return math.NaN()
		}
		res += cashflows[i] / math.Pow(1+rate, float64(math.Round(t2.Sub(t1).Hours()/24))/365)
	}
	return res + cashflows[0]
}

// Calculates the net present values of a series of cashflows at the given dates. The period length is assumed to be 1 year, accounting for leap years.
// In many cases, this returns a different result from Xnpv, which adheres to the formula used by Excel and assumes a year is always 365 days.
// cashflows and dates must be of equal length. Each cashflow is expected to occur at the date stored at the corresponding index of dates.
func Xnpv1(rate float64, cashflows []float64, dates []time.Time) float64 {
	var res float64
	if len(cashflows) != len(dates) || len(cashflows) == 0 {
		return math.NaN()
	}
	start := dates[0]
	sy, sd := start.Year(), start.YearDay()
	for i := 1; i < len(cashflows); i++ {
		res += cashflows[i] / math.Pow(1+rate, float64(dates[i].Year()-sy)+float64(dates[i].YearDay()-sd)/(365+float64(leapYear(dates[i]))))
	}
	return res + cashflows[0]
}
