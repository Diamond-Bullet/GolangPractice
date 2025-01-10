package tree

import "GolangPractice/algo/basics"

// TreeNode # binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenTree(values []int) *TreeNode {
	if values == nil {
		return nil
	}

	n := len(values)
	nodes := make([]*TreeNode, n)

	nodes[0] = &TreeNode{Val: values[0]}
	for i := 0; i <= n/2; i++ {
		if 2*i+1 < n {
			nodes[2*i+1] = &TreeNode{Val: values[2*i+1]}
			nodes[i].Left = nodes[2*i+1]
		}
		if 2*i+2 < n {
			nodes[2*i+2] = &TreeNode{Val: values[2*i+2]}
			nodes[i].Left = nodes[2*i+2]
		}
	}

	return nodes[0]
}

// InOrderWalk MarkIt 非递归,二叉树,中序遍历
func InOrderWalk(root *TreeNode) []int {
	stack := basics.Stack[*TreeNode]{}
	var ret []int
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack.Push(node)
			node = node.Left
		}
		node = stack.Pop()
		ret = append(ret, node.Val)
		node = node.Right
	}
	return ret
}

// PreOrderWalk MarkIt 非递归,二叉树,前序遍历
func PreOrderWalk(root *TreeNode) []int {
	stack := basics.Stack[*TreeNode]{}
	var ret []int
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack.Push(node)
			ret = append(ret, node.Val)
			node = node.Left
		}
		node = stack.Pop()
		node = node.Right
	}
	return ret
}

// PostOrderWalk MarkIt 非递归,二叉树,后序遍历
func PostOrderWalk(root *TreeNode) []int {
	stack := basics.Stack[*TreeNode]{}
	var ret []int
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack.Push(node)
			ret = append(ret, node.Val)
			node = node.Right
		}
		node = stack.Pop()
		node = node.Left
	}

	if ret != nil {
		n := len(ret)
		for i := 0; i < n/2; i++ {
			ret[i], ret[n-i-1] = ret[n-i-1], ret[i]
		}
	}
	return ret
}
