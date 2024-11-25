package data_structures

import "golang.org/x/exp/constraints"

// Heap
// the implement in Golang is a Min Heap.
// DO NOT use the methods of Push and Pop we write below. Import 'heap' and Use heap.Init,heap.Push,heap.Pop instead.
type Heap[T constraints.Ordered] []T

func (h *Heap[T]) Len() int           { return len(*h) }
func (h *Heap[T]) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *Heap[T]) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *Heap[T]) Push(x T) {
	*h = append(*h, x)
}

func (h *Heap[T]) Pop() (v T) {
	v, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return v
}
