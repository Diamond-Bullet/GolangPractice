package engineering

import (
	"GolangPractice/utils/logger"
	"encoding/json"
	"golang.org/x/exp/constraints"
	"testing"
)

// Generics, also named type parameters.
// equals interface plus reflect.

func TestGenericStack(t *testing.T) {
	stack := Stack[any]{}

	stack.Push(12)
	stack.Push(Stack[int]{})
	value, ok := stack.Pop()
	if !ok {
		return
	}
	logger.Info(value)

	staticStack := Stack[int]{}
	staticStack.Push(122)
	staticStack.Push(344)
	staticValue, ok := staticStack.Pop()
	if !ok {
		return
	}
	logger.Info(staticValue)
}

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

// Generic Set type
type Set[T comparable] map[T]struct{}

// Add inserts a new element into the set
func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

// Contains checks if the set contains the specified element
func (s Set[T]) Contains(value T) bool {
	_, exists := s[value]
	return exists
}

// Generic quicksort function
func QuickSort[T constraints.Ordered](data []T) []T {
	if len(data) < 2 {
		return data
	}
	left, right := 0, len(data)-1
	pivot := len(data) / 2
	data[pivot], data[right] = data[right], data[pivot]
	for i := range data {
		if data[i] < data[right] {
			data[left], data[i] = data[i], data[left]
			left++
		}
	}
	data[left], data[right] = data[right], data[left]
	QuickSort(data[:left])
	QuickSort(data[left+1:])
	return data
}

// use generics to implement more graceful result checking.
type GenericResult[T any] struct {
	Result T
	err    error
}

func (g *GenericResult[T]) Value() T {
	return g.Result
}

func (g *GenericResult[T]) Err() error {
	return g.err
}

func Validate[T any](results ...*GenericResult[T]) error {
	for _, result := range results {
		if result.Err() != nil {
			return result.Err()
		}
	}
	return nil
}

func TestGenericResult(t *testing.T) {
	b, err := json.Marshal("")
	result := &GenericResult[[]byte]{
		Result: b,
		err:    err,
	}
	if Validate(result) != nil {
		return
	}
	logger.Info(result.Value())
}
