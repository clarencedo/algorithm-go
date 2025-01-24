package leetcode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

func reverserList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre

}