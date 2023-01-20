package main

import (
	"container/list"
	"encoding/json"
	"math"
	"math/bits"
	"sort"
	"strings"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func ScatteredExamples() {
	// a-z：97-122，A-Z：65-90，0-9：48-57
}

func ForBits() {
	println(0b11111 &^ 0b00000) // 位置0

	println(bits.OnesCount(6)) // return 2. represent the count of `1` of x in binary form.

	println(bits.Len(8)) // return 4. indicate the length of x in binary form when trimming all leading zeros.
}

func ForSort() {
	// sort.Slice, sort the slice
	x := []int{3, 4, 2, 1}
	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})

	// sort.Search, f(ret) = true, f(ret-1) = false
	println(sort.Search(len(x), func(i int) bool {
		return x[i] >= 2
	}))

	// equivalent to sort.Search(len(x), func(i int) bool {
	//	return x[i] >= 2
	//})
	println(sort.SearchInts(x, 2))
}

func ForStrings() {
	strings.Fields("   ttt   eee   sss   ") // return ["ttt", "eee", "sss"]. separated by `Space` despite the count of it.

	strings.Count("sss   eee  dd ", "s")

	strings.Join([]string{"t", "dd", "s"}, "1&&1")
}

func ForRedBlackTree() {
	// RedBlackTree is also called sorted set in other languages like Java.
	// it offers the time complexity of O(log n) for both insertion and deletion.

	t := redblacktree.NewWithIntComparator() // sorted by `key`, and you can store something in `value`.
	t.Put(1, 1)
	t.Remove(1)

	t.Left()  // smallest
	t.Right() // biggest

	t.Ceiling(1) // the smallest item greater than `key`
	t.Floor(1)   // the biggest item smaller than `key`
}

func ForDoublyLinkedList() {
	println(list.New())
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Pow(a int, b float64) int {
	return int(math.Pow(float64(a), b))
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// frequently used data structure definition

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode related
func ArrayToGoSlice(arr string) interface{} {
	var x [][]int // todo: here is not dynamic
	_ = json.Unmarshal([]byte(arr), &x)
	return x
}
