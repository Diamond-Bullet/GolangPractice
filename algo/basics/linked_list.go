package basics

// ListNode linked list
type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}

// DoublyListNode doubly-linked list
type DoublyListNode[T any] struct {
	Key  T
	Val  T
	Pre  *DoublyListNode[T]
	Next *DoublyListNode[T]
}
