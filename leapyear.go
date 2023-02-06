package finlib

import (
	"time"
)

// returns 1 if t is a leapyear and 0 if it is not a leapyear
func leapYear(t time.Time) int {
	if t.Year()%4 != 0 || (t.Year()%100 == 0 && t.Year()%400 != 0) {
		return 0
	}
	return 1
}
