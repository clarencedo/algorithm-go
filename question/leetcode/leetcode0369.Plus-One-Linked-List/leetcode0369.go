package leetcode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

func plusOne(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Val: 0, Next: head}
	notNine := dummy
	//找到最后一个值不为9的节点
	for node := head; node != nil; node = node.Next {
		if node.Val != 9 {
			notNine = node
		}
	}
	//将这个节点+1
	notNine.Val++
	//将其后的所有节点变为0
	for notNine := notNine.Next; notNine != nil; notNine = notNine.Next {
		notNine.Val = 0
	}

	if dummy.Val == 1 {
		return dummy
	}
	return dummy.Next

}
