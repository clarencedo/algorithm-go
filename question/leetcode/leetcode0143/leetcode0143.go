package leetcode

import (
	"clarencedu/algorithm-go/structure"
)

type ListNode = structure.ListNode

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := findMid(head)
	left := head
	right := mid.Next
	mid.Next = nil
	right = reverseList(right)
	mergeList(left, right)
}

func findMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

func reverseListII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverseListII(head.Next)
	head.Next.Next = head
	head.Next = nil

	return last
}

func mergeList(l1, l2 *ListNode) {
	var l1Next, l2Next *ListNode
	for l1 != nil && l2 != nil {
		l1Next = l1.Next
		l2Next = l2.Next

		l1.Next = l2
		l1 = l1Next
		l2.Next = l1
		l2 = l2Next
	}
}

func reorderListTest(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	mid := findMid(head)
	left := head
	right := mid.Next
	mid.Next = nil
	right = reverseList(right)
	mergeList(left, right)
	return head
}