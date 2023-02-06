package finlib

import (
	"fmt"
	"math"
)

// Returns the alpha of s1 with respect to s2 (the benchmark) and rf (the risk free rate).
func Alpha(s1, s2 []float64, rf float64, p Precision) float64 {
	// α = R - Rf - β*(Rm - Rf)
	if s1[0] == 0 || s2[0] == 0 {
		fmt.Println(s1[0], s2[0])
		return math.NaN()
	}
	beta := Beta(s1, s2, p)
	r := s1[len(s1)-1]/s1[0] - 1  // portfolio return
	rm := s2[len(s2)-1]/s2[0] - 1 // market return
	return r - rf - beta*(rm-rf)
}
