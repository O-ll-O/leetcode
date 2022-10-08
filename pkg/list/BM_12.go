package list

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	nodes := []*ListNode{}
	curr := head
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}
	insertSortForList(nodes)
	dummy := &ListNode{}
	point := dummy
	for i := range nodes {
		point.Next = nodes[i]
		point = point.Next
	}
	return dummy.Next
}

func insertSortForList(nodes []*ListNode) {
	nodes[0].Next = nil
	for i := 1; i < len(nodes); i++ {
		temp := nodes[i]
		nodes[i].Next = nil
		j := i - 1
		for ; j >= 0; j-- {
			if nodes[j].Val > temp.Val {
				nodes[j+1] = nodes[j]
				continue
			}
			break
		}
		nodes[j+1] = temp
	}
}
