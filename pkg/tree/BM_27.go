package tree

func Print(root *TreeNode) [][]int {
	// write code here
	if root == nil {
		return [][]int{}
	}
	queue := NewQueue(1000)
	queue.Enqueue(root)
	records := make([][]int, 0)
	for !queue.IsEmpty() {
		size := queue.Size() //记录当前层多少个数
		record := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue.Dequeue() //弹出当前层的数据
			record[i] = node.Val
			//压入下一层的数据
			if node.Left != nil {
				queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}
		records = append(records, record)
	}

	for i := 0; i < len(records); i++ {
		if i%2 == 1 {
			for j, k := 0, len(records[i])-1; j < k; j, k = j+1, k-1 {
				records[i][j], records[i][k] = records[i][k], records[i][j]
			}
		}
	}
	return records
}
