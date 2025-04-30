package leetcode

import (
	"clarencedu/algorithm-go/structure"
)

type ListNode = structure.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// a,b代表k个节点的前后节点
	// 定位当前组的起始(a)和下一组的起始(b)
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

// 反转[a,b)区间的链表
func reverse(head *ListNode, target *ListNode) *ListNode {
	var pre, cur, next *ListNode
	pre, cur, next = nil, head, head
	for cur != target {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
