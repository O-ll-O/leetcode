package tree

var n, p int

func isCompleteTree(root *TreeNode) bool {
	// write code here
	if !dfsCT(root, 1) {
		return false
	}
	return n == p
}

func dfsCT(node *TreeNode, k int) bool {
	if node == nil {
		return true
	}
	n++
	if p < k {
		p = k
	}
	if k > 100 {
		return false
	}
	return dfsCT(node.Left, 2*k) && dfsCT(node.Right, 2*k+1)
}
