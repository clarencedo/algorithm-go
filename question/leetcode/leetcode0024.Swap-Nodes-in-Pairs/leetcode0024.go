package leetcode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

// recursive
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first, second := head, head.Next
	first.Next = swapPairs(second.Next)
	second.Next = first
	return second
}