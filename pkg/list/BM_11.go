package list

//https://leetcode.cn/problems/add-two-numbers-ii/submissions/
func addTwoNumbers(l3 *ListNode, l4 *ListNode) *ListNode {
	l1 := recover(l3)
	l2 := recover(l4)
	var nextCarry bool
	dummy := &ListNode{}
	curr := dummy
	for l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		v := v1 + v2
		if nextCarry {
			v++
			nextCarry = false
		}
		curr.Next = &ListNode{v, nil}
		if curr.Next.Val >= 10 {
			curr.Next.Val = v % 10
			nextCarry = true
		}
		curr = curr.Next
	}
	if nextCarry {
		curr.Next = &ListNode{1, nil}
	}
	return recover(dummy.Next)

}

func recover(pHead *ListNode) *ListNode {
	// write code here
	var pre *ListNode
	curr := pHead
	for curr != nil {
		temp := curr.Next
		curr.Next = pre
		pre = curr
		curr = temp
	}
	return pre
}
