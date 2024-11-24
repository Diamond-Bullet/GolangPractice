package tree

// A segment tree indicates some interval properties, like array.
// every node we mark as ([l, r], p), so `p` could be sum, maximum, minimum or other characteristics.

//Reference: https://oi-wiki.org/ds/seg/
// https://blog.csdn.net/weixin_42638946/article/details/115512941
// https://blog.csdn.net/andrewgithub/article/details/123435419

// Data Structure: a linked list, or an array.

// Operations:
//（1）pushUp：由子节点计算父节点的信息；
//（2）pushDown：把当前父节点的修改信息下传到子节点，也被称为懒标记（延迟标记）；这个操作比较复杂，一般不涉及到区间修改则不用写。
//（3）build：将一段区间初始化成线段树；
//（4）modify：修改操作，分为两类：① 单点修改（pushUp），② 区间修改（需要pushDown）；
//（5）query：查询一段区间的值。

// define the property of the SegmentTree.
// eg.: return i+j;
// return i*j;
// etc.
type mergeFunc func(i, j int) int

// SegmentTree not all the fields are needed, follow definition is a self-contained demonstration.
type SegmentTree struct {
	data, tree, lazy []int // data 原数据， tree 各个子节点之和
	left, right int
	merge       mergeFunc
}

// Init MarkIt 线段树 初始化
func (st *SegmentTree) Init(nums []int, op mergeFunc) {
	st.merge = op
	st.data, st.tree, st.lazy = make([]int, len(nums)), make([]int, 4*len(nums)), make([]int, 4*len(nums))
	// 将线段树中中需要存储的数据，存储到data中
	for i := 0; i < len(nums); i++ {
		st.data[i] = nums[i]
	}
	if len(nums) > 0 {
		// 构建线段树
		st.BuildSegmentTree(0, 0, len(nums)-1)
	}
}

// BuildSegmentTree MarkIt 构造线段树
func (st *SegmentTree) BuildSegmentTree(treeIndex, left, right int) {
	// 如果left = right说明已经走到了叶节点
	if left == right {
		st.tree[treeIndex] = st.data[left]
		return
	}

	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.LeftChild(treeIndex), st.RightChild(treeIndex)
	st.BuildSegmentTree(leftTreeIndex, left, midTreeIndex)
	st.BuildSegmentTree(rightTreeIndex, midTreeIndex, right)
	// 当前节点的sum值，是左 + 右
	st.tree[treeIndex] = st.merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

func (st *SegmentTree) LeftChild(index int) int {
	return index<<1 | 1 // 2x+1
}

func (st *SegmentTree) RightChild(index int) int {
	return index<<1 + 2
}

// Query MarkIt 查询 [left, right]区间的值
func (st *SegmentTree) Query(left, right int) int {
	if len(st.data) > 0 {
		return st.QueryInTree(0, 0, len(st.data)-1, left, right)
	}
	return 0
}

// QueryInTree 在一 treeIndex位根的线段树中 [left .... right]的范围内，搜索区间 [queryLeft...queryRight]的值
func (st *SegmentTree) QueryInTree(treeIndex, left, right, queryLeft, queryRight int) int {
	if left == queryLeft && right == queryRight {
		return st.tree[treeIndex]
	}

	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.LeftChild(treeIndex), st.RightChild(treeIndex)
	// 说明在树右分支
	if queryLeft > midTreeIndex {
		return st.QueryInTree(rightTreeIndex, midTreeIndex+1, right, queryLeft, queryRight)
	} else if queryRight <= midTreeIndex {
		return st.QueryInTree(leftTreeIndex, left, midTreeIndex, queryLeft, queryRight)
	}
	// 返回 左 + 右
	return st.merge(st.QueryInTree(leftTreeIndex, left, midTreeIndex, queryLeft, midTreeIndex),
		st.QueryInTree(rightTreeIndex, midTreeIndex+1, right, midTreeIndex+1, queryRight))
}

// QueryLazy MarkIt 查询 [left....right] 区间内的值
func (st *SegmentTree) QueryLazy(left, right int) int {
	if len(st.data) > 0 {
		return st.QueryLazyInTree(0, 0, len(st.data)-1, left, right)
	}
	return 0
}

// QueryLazyInTree treeIndex 查询的根节点，[left,right] 被查询区间， [queryLeft, queryRight]需要去查询的区间
func (st *SegmentTree) QueryLazyInTree(treeIndex, left, right, queryLeft, queryRight int) int {
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.LeftChild(treeIndex), st.RightChild(treeIndex)
	// 看需要查询的区间是否在被查询区间内
	if left > queryRight || right < queryLeft {
		// 如果查询位置超越实际区间，直接放回0
		return 0 // represents a null node
	}
	// 当对应节点的lazy值非空，说明该节点使用了懒查询和懒更新
	if st.lazy[treeIndex] != 0 { // this node is lazy
		for i := 0; i < right-left+1; i++ {
			// 当前tree节点值，+ 所有子节点的lazy和
			st.tree[treeIndex] = st.merge(st.tree[treeIndex], st.lazy[treeIndex])
		}
		// 将lazy值想子节点传递
		if left != right { // update lazy[] for children nodes
			st.lazy[leftTreeIndex] = st.merge(st.lazy[leftTreeIndex], st.lazy[treeIndex])
			st.lazy[rightTreeIndex] = st.merge(st.lazy[rightTreeIndex], st.lazy[treeIndex])
		}
		// 处理完成之后，当前节点不再是懒节点
		st.lazy[treeIndex] = 0 // current node processed. No longer lazy
	}
	// 如果当前区间在查询区间的内部，直接放回当前区间的tree值即可
	if queryLeft <= left && queryRight >= right { // segment completely inside range
		return st.tree[treeIndex]
	}
	// 全部在右半部分
	if queryLeft > midTreeIndex {
		return st.QueryLazyInTree(rightTreeIndex, midTreeIndex+1, right, queryLeft, queryRight)
	} else if queryRight <= midTreeIndex {
		// 全部在右半部分
		return st.QueryLazyInTree(leftTreeIndex, left, midTreeIndex, queryLeft, queryRight)
	}
	// 如果中线两边都有，则按照中线  进行 左 + 右
	// merge query results
	return st.merge(st.QueryLazyInTree(leftTreeIndex, left, midTreeIndex, queryLeft, midTreeIndex),
		st.QueryLazyInTree(rightTreeIndex, midTreeIndex+1, right, midTreeIndex+1, queryRight))
}

// Update MarkIt 更新树节点
func (st *SegmentTree) Update(index, val int) {
	if len(st.data) > 0 {
		st.UpdateInTree(0, 0, len(st.data)-1, index, val)
	}
}

// UpdateInTree 以 treeIndex 为根，更新 index 位置上的值为 val
func (st *SegmentTree) UpdateInTree(treeIndex, left, right, index, val int) {
	// 找到节点
	if left == right {
		st.tree[treeIndex] = val
		return
	}
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.LeftChild(treeIndex), st.RightChild(treeIndex)
	// 如果index在右子树
	if index > midTreeIndex {
		st.UpdateInTree(rightTreeIndex, midTreeIndex+1, right, index, val)
	} else {
		// 如果index在左子树
		st.UpdateInTree(leftTreeIndex, left, midTreeIndex, index, val)
	}
	// 退出程序之前将所有涉及到的节点值都更新下
	st.tree[treeIndex] = st.merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

// UpdateLazy 区间更新函数定义
func (st *SegmentTree) UpdateLazy(updateLeft, updateRight, val int) {
	if len(st.data) > 0 {
		st.UpdateLazyInTree(0, 0, len(st.data)-1, updateLeft, updateRight, val)
	}
}

func (st *SegmentTree) UpdateLazyInTree(treeIndex, left, right, queryLeft, queryRight, val int) int {
	return 0
}
