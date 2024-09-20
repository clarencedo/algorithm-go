package leetcode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	targetNode := findNthFromEnd(dummy, n+1)
	targetNode.Next = targetNode.Next.Next
	return dummy.Next
}

func findNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}