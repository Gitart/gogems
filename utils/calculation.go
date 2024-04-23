package utils

type Number interface {
	int16 | int32 | int64 | float32 | float64
}

// optional types in generic functions
func SumAny[T int64 | int | float64](t ...T) *T {
	var s T

	for _, ts := range t {
		s += ts
	}

	return &s
}
