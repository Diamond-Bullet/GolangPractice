package utils

import (
	"golang.org/x/exp/constraints"
	"math"
)

func Ptr[T any](v T) *T {
	return &v
}

type Number interface {
	constraints.Integer | constraints.Float
}

func Pow[T Number](x, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
