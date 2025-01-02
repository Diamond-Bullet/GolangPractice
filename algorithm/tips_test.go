package main

import (
	"GolangPractice/pkg/logger"
	"container/list"
	"github.com/emirpasic/gods/trees/redblacktree"
	"math/bits"
	"sort"
	"strings"
	"testing"
)

func TestScattered(t *testing.T) {
	// a-z：97-122，A-Z：65-90，0-9：48-57
}

func TestBits(t *testing.T) {
	// https://leetcode.cn/problems/insert-into-bits-lcci/solutions/2458770/jian-dan-wei-yun-suan-by-raccooncc-l7jo/?envType=study-plan-v2&envId=cracking-the-coding-interview

	// set to 0 by bit.
	// z = x &^ y. if the bit in `y` is 1, then the corresponding bit in `z` is 0.
	logger.Infof("%b", 0b11111&^0b00001)
	// represent the count of `1` of x in binary form.
	logger.Info(bits.OnesCount(6)) // 2
	// indicate the length of x in binary form when trimming all leading zeros.
	logger.Info(bits.Len(8)) // 4
}

func TestSort(t *testing.T) {
	// sort.Slice, sort the slice
	x := []int{3, 4, 2, 1}
	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})

	// sort.Search, f(ret) = true, f(ret-1) = false
	logger.Info(sort.Search(len(x), func(i int) bool {
		return x[i] >= 2
	}))

	// equivalent to sort.Search(len(x), func(i int) bool {
	//	return x[i] >= 2
	//})
	logger.Info(sort.SearchInts(x, 2))
}

func TestStrings(t *testing.T) {
	strings.Fields("   ttt   eee   sss   ") // return ["ttt", "eee", "sss"]. separated by `Space` despite the count of it.

	strings.Count("sss   eee  dd ", "s")

	strings.Join([]string{"t", "dd", "s"}, "1&&1")
}

func TestRedBlackTree(t *testing.T) {
	// RedBlackTree is also called sorted set in other languages like Java.
	// it offers the time complexity of O(log n) for both insertion and deletion.

	tree := redblacktree.NewWithIntComparator() // sorted by `key`, and you can store something in `value`.
	tree.Put(1, 1)
	tree.Remove(1)

	tree.Left()  // smallest
	tree.Right() // biggest

	tree.Ceiling(1) // the smallest item greater than `key`
	tree.Floor(1)   // the biggest item smaller than `key`
}

func TestDoublyLinkedList(t *testing.T) {
	logger.Info(list.New())
}
