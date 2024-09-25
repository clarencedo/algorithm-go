package leetcode

import "clarencedu/algorithm-go/structure"

type Node = structure.ListNode

func insert(aNode *Node, x int) *Node {
	node := &Node{Val: x}
	if aNode == nil {
		node.Next = node
		return node
	}
	if aNode.Next == aNode {
		aNode.Next = node
		node.Next = aNode
		return aNode
	}

	cur, next := aNode, aNode.Next
	for next != aNode {
		if x >= cur.Val && x <= next.Val {
			break
		}
		// 指针开始的位置 不一定是链表头，可能是链表中间的某个位置
		if cur.Val > next.Val && (x > cur.Val || x < next.Val) {
			break
		}
		cur, next = next, next.Next
	}
	cur.Next = &Node{Val: x, Next: next}

	return aNode
}