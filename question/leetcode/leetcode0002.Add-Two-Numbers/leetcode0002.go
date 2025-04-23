package leetcode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

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

// 链表求和,正序存储
func addTwoNumbersForward(l1 *ListNode, l2 *ListNode) *ListNode {
	// 反转链表
	l1 = reverseList(l1)
	l2 = reverseList(l2)

	// 逆序相加
	result := addTwoNumbers(l1, l2)

	// 反转结果链表
	return reverseList(result)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

// 栈+ 递归
func addTwoNumbersForwardII(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := []int{}, []int{}
	// 将链表的值压入栈
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	carry := 0
	var result *ListNode
	// 从栈顶开始相加
	for len(stack1) > 0 || len(stack2) > 0 || carry != 0 {
		n1, n2 := 0, 0
		if len(stack1) > 0 {
			n1 = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			n2 = stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}

		sum := n1 + n2 + carry
		carry = sum / 10
		newNode := &ListNode{Val: sum % 10}
		newNode.Next = result
		result = newNode
	}

	return result
}