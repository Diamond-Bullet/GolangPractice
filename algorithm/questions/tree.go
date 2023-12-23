package questions

import (
	"GolangPractice/algorithm/data_structures"
)

// InOrderWalk MarkIt 非递归,二叉树,中序遍历
func InOrderWalk(root *data_structures.TreeNode) []int {
	stack := data_structures.StackTemplate{}
	var ret []int
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack.Push(node)
			node = node.Left
		}
		node = stack.Pop().(*data_structures.TreeNode)
		ret = append(ret, node.Val)
		node = node.Right
	}
	return ret
}

// PreOrderWalk MarkIt 非递归,二叉树,前序遍历
func PreOrderWalk(root *data_structures.TreeNode) []int {
	stack := data_structures.StackTemplate{}
	var ret []int
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack.Push(node)
			ret = append(ret, node.Val)
			node = node.Left
		}
		node = stack.Pop().(*data_structures.TreeNode)
		node = node.Right
	}
	return ret
}

// PostOrderWalk MarkIt 非递归,二叉树,后序遍历
func PostOrderWalk(root *data_structures.TreeNode) []int {
	stack := data_structures.StackTemplate{}
	var ret []int
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack.Push(node)
			ret = append(ret, node.Val)
			node = node.Right
		}
		node = stack.Pop().(*data_structures.TreeNode)
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

// TreeDelete MarkIt 二叉搜索树，删除节点
func TreeDelete(root, target *data_structures.TreeNode) {
	targetParent := NodeParent(root, target)
	if target.Left == nil {
		TransPlant(root, target, target.Right, targetParent, target)
	} else if target.Right == nil {
		TransPlant(root, target, target.Left, targetParent, target)
	} else {
		y := TreeMinimum(target.Right)
		yParent := NodeParent(root, y)
		if yParent != target {
			TransPlant(root, y, y.Right, yParent, y)
			y.Right = target.Right
			target.Right = nil
		}
		TransPlant(root, target, y, targetParent, yParent)
		y.Left = target.Left
		target.Left = nil
	}
}

func TransPlant(root, u, v, uParent, vParent *data_structures.TreeNode) {
	if uParent == nil {
		root = v
	} else if u == uParent.Left {
		uParent.Left = v
	} else {
		uParent.Right = v
	}
	if v != nil {
		vParent = uParent
	}
}

func NodeParent(root, target *data_structures.TreeNode) *data_structures.TreeNode {
	node := root
	for node != target {
		if node.Val > target.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return node
}

func TreeMinimum(root *data_structures.TreeNode) *data_structures.TreeNode {
	if root == nil {
		return nil
	}
	node := root
	for node.Left != nil {
		node = node.Left
	}
	return node
}
