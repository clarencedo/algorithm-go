package leetocode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

// 两
func addTwoNumbers445(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseListII(l1)
	l2 = reverseListII(l2)
	l3 := addTwoNumbers(l1, l2)

	return reverseListII(l3)
}

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	n1, n2, carry, cur := 0, 0, 0, dummy
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		cur = cur.Next
		carry = (n1 + n2 + carry) / 10
	}
	return dummy.Next
}