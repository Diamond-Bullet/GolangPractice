package data_structures

// ListNode linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// DoublyListNode doubly-linked list
type DoublyListNode struct {
	Key  int
	Val  int
	Pre  *DoublyListNode
	Next *DoublyListNode
}

