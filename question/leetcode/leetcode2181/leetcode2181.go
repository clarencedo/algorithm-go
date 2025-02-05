package leetcode

import (
	"clarencedu/algorithm-go/structure"
)

type ListNode = structure.ListNode

func mergeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{}
	list := dummy
	sum := 0

	for cur := head.Next; cur != nil; cur = cur.Next {
		if cur.Val == 0 {
			node := &ListNode{Val: sum}
			list.Next = node
			list = list.Next
			sum = 0
		} else {
			sum += cur.Val
		}
	}

	return dummy.Next
}