package list

//https://www.nowcoder.com/practice/d8b6b4358f774294a89de2a6ac4d9337?tpId=295&tqId=23267&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
func Merge(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{0, nil}
	pre := prehead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			pre = pre.Next
			l1 = l1.Next
			continue
		}
		pre.Next = l2
		pre = pre.Next
		l2 = l2.Next

	}
	if l1 != nil {
		pre.Next = l1
	}
	if l2 != nil {
		pre.Next = l2
	}
	return prehead.Next
}
