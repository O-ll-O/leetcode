package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param root TreeNode类
 * @return int整型一维数组
 */
func preorderTraversal(root *TreeNode) []int {
	// write code here
	data := &[]int{}
	preTraversal(root, data)
	return *data
}

func preTraversal(node *TreeNode, data *[]int) {
	if node == nil {
		return
	}
	*data = append(*data, node.Val)
	preTraversal(node.Left, data)
	preTraversal(node.Right, data)

}

func inOrder(node *TreeNode, data *[]int) {
	if node == nil {
		return
	}
	inOrder(node.Left, data)
	*data = append(*data, node.Val)
	inOrder(node.Right, data)

}

func inPost(node *TreeNode, data *[]int) {
	if node == nil {
		return
	}
	inPost(node.Left, data)
	inPost(node.Right, data)
	*data = append(*data, node.Val)
}
