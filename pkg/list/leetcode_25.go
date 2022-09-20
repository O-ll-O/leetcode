package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy
	end := dummy
	for end.Next != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		start := pre.Next
		next := end.Next
		end.Next = nil
		pre.Next = reverse(start)
		start.Next = next
		pre = start
		end = pre

	}
	return dummy.Next
}

func reverse(pHead *ListNode) *ListNode {
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
