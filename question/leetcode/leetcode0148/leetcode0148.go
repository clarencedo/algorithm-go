package leetcode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

func sortList(head *ListNode) *ListNode {
	return mergeSort(head, nil)
}

// 归并
func mergeSort(head, tail *ListNode) *ListNode {

	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast.Next != tail {
			fast = fast.Next
		}
	}
	mid := slow

	return mergeTwoList(mergeSort(head, mid), mergeSort(mid, tail))

}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next, l1 = l1, l1.Next
		} else {
			current.Next, l2 = l2, l2.Next
		}
	}

	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}

	return dummy.Next
}