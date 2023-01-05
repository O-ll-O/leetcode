package tree

func IsBalanced_Solution(pRoot *TreeNode) bool {
	// write code here
	return hight(pRoot) > 0
}

func hight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := hight(node.Left)
	right := hight(node.Right)
	if left == -1 || right == -1 || abs(left-right) > 1 {
		return -1
	}
	return max(left, right) + 1
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
