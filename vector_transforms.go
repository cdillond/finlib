package finlib

import (
	"math"
)

func Acos(s []float64) []float64 {
	res := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = math.Acos(s[i])
	}
	return res
}

func Asin(s []float64) []float64 {
	res := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = math.Asin(s[i])
	}
	return res
}

func Atan(s []float64) []float64 {
	res := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = math.Atan(s[i])
	}
	return res
}

func Ceil(s []float64) []float64 {
	res := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = math.Ceil(s[i])
	}
	return res
}

func Cos(s []float64) []float64 {
	res := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = math.Cos(s[i])
	}
	return res
}

func Cosh(s []float64) []float64 {
	res := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = math.Cosh(s[i])
	}
	return res
}

func Exp(s []float64) []float64 {
	res := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = math.Exp(s[i])
	}
	return res
}

func Floor(s []float64) []float64 {
	res := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = math.Floor(s[i])
	}
	return res
}

// as LN in python documentation
func Log(s []float64) []float64 {
	res := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = math.Log(s[i])
	}
	return res
}

func Log10(s []float64) []float64 {
	res := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = math.Log10(s[i])
	}
	return res
}

func Sin(s []float64) []float64 {
	res := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = math.Sin(s[i])
	}
	return res
}

func Sinh(s []float64) []float64 {
	res := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = math.Sinh(s[i])
	}
	return res
}

func Sqrt(s []float64) []float64 {
	res := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = math.Sqrt(s[i])
	}
	return res
}

func Tan(s []float64) []float64 {
	res := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = math.Tan(s[i])
	}
	return res
}

func Tanh(s []float64) []float64 {
	res := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = math.Tanh(s[i])
	}
	return res
}
