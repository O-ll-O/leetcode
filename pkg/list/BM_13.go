package list

func isPalindrome(head *ListNode) bool {
	curr := head
	ints := []int{}
	for curr != nil {
		ints = append(ints, curr.Val)
		curr = curr.Next
	}
	for i, j := 0, len(ints)-1; j > i; i, j = i+1, j-1 {
		if ints[i] != ints[j] {
			return false
		}
	}
	return true
}

func recover2(pHead *ListNode) *ListNode {
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
