package leetcode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

// INFO: 统计链表的长度，然后分割链表, n/k 确定每个部分长度, n%k 确定前面几个部分长度+1
// INFO: partSize 代表当前部分链表的节点数，curr当前链表节点指针,每拆分一部分后，将链表断开 curr.Next = nil
func splitListToParts(head *ListNode, k int) []*ListNode {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}

	quotient, remainder := n/k, n%k
	parts := make([]*ListNode, k)

	for i, curr := 0, head; i < k && curr != nil; i++ {
		parts[i] = curr
		partSize := quotient

		if i < remainder {
			partSize++
		}

		for j := 1; j < partSize; j++ {
			curr = curr.Next
		}

		//curr 指向下一个链表的头节点,并且断开连接
		curr, curr.Next = curr.Next, nil
	}

	return parts
}