package miscellaneous

import (
	"GolangPractice/algorithm/basics/tree"
)

// TreeDelete MarkIt 二叉搜索树，删除节点
func TreeDelete(root, target *tree.TreeNode) {
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

func TransPlant(root, u, v, uParent, vParent *tree.TreeNode) {
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

func NodeParent(root, target *tree.TreeNode) *tree.TreeNode {
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

func TreeMinimum(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}
	node := root
	for node.Left != nil {
		node = node.Left
	}
	return node
}
