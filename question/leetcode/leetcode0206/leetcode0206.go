package leetcode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

// 方法一：迭代
func reverserList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre

}

// 方法二：递归
func reverseListII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseListII(head.Next)
	head.Next.Next = head
	head.Next = nil

	return last
}
