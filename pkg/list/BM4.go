package list

//https://www.nowcoder.com/practice/d8b6b4358f774294a89de2a6ac4d9337?tpId=295&tqId=23267&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{-1, nil}
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

//https://www.nowcoder.com/practice/65cfde9e5b9b4cf2b6bafa5f3ef33fa6?tpId=295&tqId=724&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
func mergeKGroup(list []*ListNode) *ListNode {
	var temp *ListNode = nil
	for i := 0; i < len(list)-1; i++ {
		temp = mergeTwoList(temp, list[i])
	}
	return temp
}

//https://leetcode.cn/problems/merge-k-sorted-lists/
func mergeKLists(list []*ListNode) *ListNode {
	return mergeList(list, len(list))
}

func mergeList(list []*ListNode, n int) *ListNode {
	return mergetList_c(list, 0, n-1)
}

func mergetList_c(list []*ListNode, p, r int) *ListNode {
	if p == r {
		return list[p]
	}
	if p > r {
		return nil
	}
	mid := (p + r) >> 1
	return mergeTwoList(mergetList_c(list, p, mid), mergetList_c(list, mid+1, r))
}

// 递推公式：
// merge_sort(p…r) = merge(merge_sort(p…q), merge_sort(q+1…r))

// 终止条件：
// p >= r 不用再继续分解

// // 归并排序算法, A是数组，n表示数组大小
// merge_sort(A, n) {
// 	merge_sort_c(A, 0, n-1)
//   }

//   // 递归调用函数
//   merge_sort_c(A, p, r) {
// 	// 递归终止条件
// 	if p >= r  then return

// 	// 取p到r之间的中间位置q
// 	q = (p+r) / 2
// 	// 分治递归
// 	merge_sort_c(A, p, q)
// 	merge_sort_c(A, q+1, r)
// 	// 将A[p...q]和A[q+1...r]合并为A[p...r]
// 	merge(A[p...r], A[p...q], A[q+1...r])
//   }

// merge(A[p...r], A[p...q], A[q+1...r]) {
// 	var i := p，j := q+1，k := 0 // 初始化变量i, j, k
// 	var tmp := new array[0...r-p] // 申请一个大小跟A[p...r]一样的临时数组
// 	while i<=q AND j<=r do {
// 	  if A[i] <= A[j] {
// 		tmp[k++] = A[i++] // i++等于i:=i+1
// 	  } else {
// 		tmp[k++] = A[j++]
// 	  }
// 	}

// 	// 判断哪个子数组中有剩余的数据
// 	var start := i，end := q
// 	if j<=r then start := j, end:=r

// 	// 将剩余的数据拷贝到临时数组tmp
// 	while start <= end do {
// 	  tmp[k++] = A[start++]
// 	}

// 	// 将tmp中的数组拷贝回A[p...r]
// 	for i:=0 to r-p do {
// 	  A[p+i] = tmp[i]
// 	}
//   }
