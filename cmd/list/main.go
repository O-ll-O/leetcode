package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, l, r int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	pre := dummy
	for i := 0; i < l-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	var temp *ListNode
	for i := l; i < r; i++ {
		temp = cur.Next
		cur.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp
	}
	return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		temp := curr.Next
		curr.Next = pre
		pre = curr
		curr = temp
	}
	return pre
}
func kthToLast()

// int kthToLast(struct ListNode* head, int k){
//     struct ListNode* first = head;
//     struct ListNode* last = head;
//     for(int i = 1; i < k ;i++)
//         first = first->Next;
//     while(first->Next)
//     {
//         first = first->Next;
//         last=last->Next;
//     }
//     return last->Val;
// }
