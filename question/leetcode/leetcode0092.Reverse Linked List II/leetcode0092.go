package leetcode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

func reverseBwtween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	head.Next = reverseBwtween(head.Next, left-1, right-1)
	return head
}

func reverseN(head *ListNode, n int) *ListNode {
	cur := head
	prev := &ListNode{Val: 0, Next: head}
	for i := 0; i < n; i++ {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	head.Next = cur
	return prev
}