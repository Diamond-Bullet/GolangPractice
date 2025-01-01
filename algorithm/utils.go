package main

import (
	"golang.org/x/exp/constraints"
	"math"
)

// some duplicate code here, for coding convenience.

type Number interface {
	constraints.Integer | constraints.Float
}

func Pow[T Number](x, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}

func Abs[T Number](x T) T {
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func List(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func LinkedList(arr []int) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	for _, v := range arr {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummyHead.Next
}
