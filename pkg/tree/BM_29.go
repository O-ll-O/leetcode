package tree

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * @param root TreeNode类
 * @param sum int整型
 * @return bool布尔型
 */
var path = []int{}
var ans = [][]int{}

func hasPathSum(root *TreeNode, sum int) bool {
	// write code here
	dfs(root, sum)
	return len(ans) != 0
}

func dfs(node *TreeNode, target int) {
	if node == nil {
		return
	}
	target -= node.Val
	path = append(path, node.Val)
	defer func() { path = path[:len(path)-1] }()
	if node.Left == nil && node.Right == nil && target == 0 {
		ans = append(ans, append([]int(nil), path...))
		return
	}
	dfs(node.Left, target)
	dfs(node.Right, target)
}
