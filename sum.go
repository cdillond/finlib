package finlib

// Returns the sum of all values in s.
func Sum(s []float64) float64 {
	return kahanSum(s)
}

// error correcting
func kahanSum(s []float64) float64 {
	var sum float64
	var c float64 // correction
	for i := 0; i < len(s); i++ {
		y := s[i] - c   // adjust s[i] based on previous error
		t := sum + y    // add corrected s[i] to the sum to get a new sum
		c = t - sum - y // subtract old sum and y from the new sum to get the error
		sum = t         // assign new sum to old sum variable
	}
	return sum
}
