package compare

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
	if a > b {
		return b
	}
	return a
}