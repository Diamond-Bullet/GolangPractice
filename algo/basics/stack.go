package basics

type Stack[T any] []T

func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) Pop() (v T) {
	theStack := *s
	if len(theStack) == 0 {
		return
	}
	*s, v = theStack[:len(theStack)-1], theStack[len(theStack)-1]
	return v
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

type StackWrapper[T any] struct {
	stack *[]T
}

func (s *StackWrapper[T]) Push(value T) {
	*s.stack = append(*s.stack, value)
}

func (s *StackWrapper[T]) Pop() (v T) {
	theStack := *s.stack
	if len(theStack) == 0 {
		return
	}
	*s.stack, v = theStack[:len(theStack)-1], theStack[len(theStack)-1]
	return v
}

func (s *StackWrapper[T]) Len() int {
	return len(*s.stack)
}
