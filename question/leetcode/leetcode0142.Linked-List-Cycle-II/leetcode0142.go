package leetcode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	hasCircle := false
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			hasCircle = true
			break
		}
	}
	if !hasCircle {
		return nil
	}
	slow = head
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}