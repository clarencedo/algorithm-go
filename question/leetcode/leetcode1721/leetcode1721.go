package leetcode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

func swapNodes(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	pre := dummy
	for i := 1; i <= k; i++ {
		pre = pre.Next
	}
	first := pre
	second := dummy.Next

	// 两个指针同时移动，当pre到达末尾时，second指向的是倒数第k个节点
	for pre != nil && pre.Next != nil {
		pre = pre.Next
		second = second.Next
	}
	first.Val, second.Val = second.Val, first.Val
	return dummy.Next

}