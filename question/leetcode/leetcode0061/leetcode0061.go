package leetcode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	n, tail := 1, head
	// 计算链表长度
	for tail.Next != nil {
		tail = tail.Next
		n++
	}

	// 首尾相连
	k = k % n
	// 实际要处理的旋转次数
	if k == 0 {
		return head
	}

	//形成环
	tail.Next = head
	newHeadIndex := n - k
	newTailNode := tail

	for i := 0; i < newHeadIndex; i++ {
		newTailNode = newTailNode.Next
	}

	newHead := newTailNode.Next
	// 断开环
	newTailNode.Next = nil

	return newHead
}
