package main

import (
	"golang.org/x/exp/constraints"
	"math"
)

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

// frequently used data structures

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}
