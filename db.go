package finlib

import "math"

// Returns the depreciation of an asset over its life using the fixed-declining balance method.
// The output of this function is intended to match the output of Excel's DB function; for greater precision, use Xdb.
// cost is the initial cost of the asset, salvage is the value at the end of the depreciation, life is the number of periods over which the asset is being depreciated, and month is the number of months in the first year.
func Db(cost, salvage, month float64, life, period uint) float64 {
	rate := 1 - math.Pow(salvage/cost, 1/float64(life))
	rate = math.Round(rate*1_000) / 1_000
	total := cost * rate * month / 12
	curr := total
	// may be numerically unstable
	for i := 1; i < int(period); i++ {
		if i == int(life) {
			curr = (cost - total) * rate * (12 - month) / 12
			break
		}
		curr = (cost - total) * rate
		total += curr
	}
	return math.Round(curr*100) / 100
}

// Returns the depreciation of an asset over its life using the fixed-declining balance method.
// cost is the initial cost of the asset, salvage is the value at the end of the depreciation, life is the number of periods over which the asset is being depreciated, and month is the number of months in the first year.
func Xdb(cost, salvage, month float64, life, period uint) float64 {
	rate := 1 - math.Pow(salvage/cost, 1/float64(life))
	total := cost * rate * month / 12
	curr := total
	// may be numerically unstable
	for i := 1; i < int(period); i++ {
		if i == int(life) {
			curr = (cost - total) * rate * (12 - month) / 12
			break
		}
		curr = (cost - total) * rate
		total += curr
	}
	return curr
}
