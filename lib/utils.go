package lib

import (
	"GolangPractice/lib/basic"
	"math"
)

func Ptr[T any](v T) *T {
	return &v
}

func Pow[T basic.Number](x, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}

func Abs[T basic.Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
