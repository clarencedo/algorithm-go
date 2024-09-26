package leetcode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

func plusOne(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Val: 0, Next: head}
	notNine := dummy
	for node := head; node != nil; node = node.Next {
		if node.Val != 9 {
			notNine = node
		}
	}
	notNine.Val++
	for node := notNine.Next; node != nil; node = node.Next {
		node.Val = 0
	}

	if dummy.Val == 0 {
		return dummy.Next
	}
	return dummy

}