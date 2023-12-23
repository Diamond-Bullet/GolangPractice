package data_structures

// StackTemplate
// 但通常你不需要这样使用，因为有额外的类型断言开销
type StackTemplate []interface{}

func (stack *StackTemplate) Push(value interface{}) {
	*stack = append(*stack, value)
}

func (stack *StackTemplate) Pop() (v interface{}) {
	theStack := *stack
	if len(theStack) == 0 {
		return
	}
	*stack, v = theStack[:len(theStack)-1], theStack[len(theStack)-1]
	return v
}

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
