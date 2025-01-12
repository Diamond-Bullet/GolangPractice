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

func Map[T, F any](source []T, mapper func(T) F) []F {
	result := make([]F, len(source))
	for i, v := range source {
		result[i] = mapper(v)
	}
	return result
}

func Reduce[T, F any](source []T, reducer func(F, T) F, initial F) F {
	result := initial
	for _, v := range source {
		result = reducer(result, v)
	}
	return result
}

func Filter[T any](source []T, filter func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range source {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BinaryTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, len(arr))
	nodes[0] = &TreeNode{Val: arr[0]}
	for i := 0; i < len(arr)>>1; i++ {
		node := nodes[i]
		if i<<1+1 < len(arr) {
			node.Left = &TreeNode{Val: arr[i<<1+1]}
			nodes[i<<1+1] = node.Left
		}
		if i<<1+2 < len(arr) {
			node.Right = &TreeNode{Val: arr[i<<1+2]}
			nodes[i<<1+2] = node.Right
		}
	}
	return nodes[0]
}
