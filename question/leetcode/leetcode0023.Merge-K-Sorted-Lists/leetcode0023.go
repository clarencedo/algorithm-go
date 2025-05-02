package leetcode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n < 1 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	mid := n / 2
	left := mergeKLists(lists[0:mid])
	right := mergeKLists(lists[mid:])
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	} else {
		l2.Next = merge(l1, l2.Next)
		return l2
	}
}
