package pkg

import (
	"GolangPractice/pkg/basic"
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

func Ternary[T any](condition bool, forTrue, forFalse T) T {
	if condition {
		return forTrue
	}
	return forFalse
}

func Keys[T comparable, V any](m map[T]V) []T {
	s := make([]T, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

func Values[T comparable, V any](m map[T]V) []V {
	s := make([]V, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}
