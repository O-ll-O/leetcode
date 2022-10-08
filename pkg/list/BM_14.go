package list

func oddEvenList(head *ListNode) *ListNode {
	// write code here
	if head == nil {
		return nil
	}
	odd := head
	even := head.Next
	eventHead := even
	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		odd = odd.Next
		even.Next = even.Next.Next
		even = even.Next

	}
	odd.Next = eventHead
	return odd
}
