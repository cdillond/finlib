package finlib

import (
	"math"
)

// Returns rounds f to the nearest cent and returns it as a string formatted as #,###.## or, if negative, (#,###.##).
func FloatToMoney(f float64) string {
	var res string
	var neg bool
	if f < 0 {
		neg = true
		f = -1 * f
	}
	f = math.Round(f * 100)
	n := int(f)

	for i := 0; i < 2; i++ {
		res = string('0'+byte(n%10)) + res
		n /= 10
	}
	res = "." + res
	if n == 0 {
		res = "0" + res
	}
	for i := 1; n > 0; n, i = n/10, i+1 {
		if i%4 == 0 {
			res = "," + res
		}
		res = string('0'+byte(n%10)) + res
	}

	if neg {
		res = "(" + res + ")"
	}
	return res
}
