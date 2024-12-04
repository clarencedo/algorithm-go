package leetcode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

func deleteNodes(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	tmp := head

	for tmp != nil && tmp.Next != nil {
		for i := 0; i < m-1 && tmp != nil; i++ {
			tmp = tmp.Next
		}

		for i := 0; i < n && tmp != nil && tmp.Next != nil; i++ {
			tmp.Next = tmp.Next.Next
		}

	}
	return dummy.Next

}