package golang

import (
	"GolangPractice/utils/logger"
	"golang.org/x/exp/constraints"
	"testing"
)

// see examples in Go official package "slices"
func Contain[T comparable](s []T, p T) bool {
	for _, ss := range s {
		if p == ss {
			return true
		}
	}
	return false
}

// define a generic type
type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Max[T Number](a, b T) T {
	if a < b {
		return b
	}
	return a
}

// package "constraints" provides commonly used types of generics.
func Min[T constraints.Ordered](a, b T) T {
	if a > b {
		return b
	}
	return a
}

// a generic type can not be used as an ordinary type.
// but it can be the one of an ordinary type's field.
type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(v T) {
	s.elements = append(s.elements, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zeroVal T // default value for the type T
		return zeroVal, false
	}
	lastIndex := len(s.elements) - 1
	v := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return v, true
}

func TestGenericStack(t *testing.T) {
	stack := Stack[any]{}

	stack.Push(12)
	stack.Push(Stack[int]{})
	value, ok := stack.Pop()
	if !ok {
		return
	}
	logger.Infoln(value)

	staticStack := Stack[int]{}
	staticStack.Push(122)
	staticStack.Push(344)
	staticValue, ok := staticStack.Pop()
	if !ok {
		return
	}
	logger.Infoln(staticValue)
}
