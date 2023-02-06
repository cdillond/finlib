package finlib

type Precision int

const (
	Default = Precision(iota) // Error corrected
	Naive
	Exact // Not recommended
	Fast  // Not recommended
)
