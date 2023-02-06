package finlib

import "math"

// Returns the depreciation of an asset over its life using the fixed-declining balance method.
// The rate is rounded to three decimal places. For greater precision, use Xdb.
func Db(cost, salvage, month float64, life uint) []float64 {
	res := make([]float64, life)
	if life == 0 {
		return res
	}
	rate := 1 - math.Pow(salvage/cost, 1/float64(life))
	rate = math.Round(rate*1_000) / 1_000
	res[0] = cost * rate * month / 12
	total := res[0]
	for i := 1; i < int(life)-1; i++ {
		n := (cost - total) * rate
		res[i] = n
		total += n
	}
	res[life-1] = (cost - total) * rate * (12 - month) / 12
	return res
}

// Returns the depreciation of an asset over its life using the fixed-declining balance method.
func Xdb(cost, salvage, month float64, life uint) []float64 {
	res := make([]float64, life)
	if life == 0 {
		return res
	}
	rate := 1 - math.Pow(salvage/cost, 1/float64(life))
	res[0] = cost * rate * month / 12
	total := res[0]
	for i := 1; i < int(life)-1; i++ {
		n := (cost - total) * rate
		res[i] = n
		total += n
	}
	res[life-1] = (cost - total) * rate * (12 - month) / 12
	return res
}

// Returns the depreciation of an asset for a specified period using the declining balance method with the specified factor.
func Fdb(cost, salvage, factor float64, period, life uint) float64 {
	//res := make([]float64, life)
	var total float64
	for i := 0; i < int(period); i++ {
		n1 := (cost - total) * (factor / float64(life))
		n2 := cost - salvage - total
		if n1 < n2 {
			//res[i] = n1
			total += n1
		} else {
			//res[i] = n2
			total += n2
		}
	}
	return total
}
