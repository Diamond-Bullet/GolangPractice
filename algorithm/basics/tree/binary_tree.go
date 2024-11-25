package tree

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
