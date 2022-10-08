package list

import (
	"fmt"
	"testing"
)

var l1, l2 *ListNode

func init() {
	l1 = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  7,
				Next: nil,
			},
		},
	}
	l2 = &ListNode{
		Val: 6,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	}
}

func TestAddTwoNumbers(t *testing.T) {
	l := addTwoNumbers(l1, l2)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func TestSortList(t *testing.T) {
	l := sortList(l1)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
