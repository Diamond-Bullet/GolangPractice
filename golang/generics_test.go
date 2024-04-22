package golang

import "golang.org/x/exp/constraints"

func Contain[T comparable](s []T, p T) bool {
	for _, ss := range s {
		if p == ss {
			return true
		}
	}
	return false
}

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Max[T Number](a, b T) T {
	if a < b {
		return b
	}
	return a
}

func Min[T constraints.Ordered](a, b T) T {
	if a > b {
		return b
	}
	return a
}