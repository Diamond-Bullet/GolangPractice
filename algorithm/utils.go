package main

import (
	"GolangPractice/lib/basic"
	"math"
)

func Pow[T basic.Number](x, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}

func Abs[T basic.Number](x T) T {
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
