package leetcode

import "clarencedu/algorithm-go/structure"

type Node = structure.RandListNode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	//在原节点后面插入一个复制节点
	// A -> B -> C
	// A -> A` -> B ->B` ->C ->C`
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}
	//复制Randow指针
	//让复制节点A'的Random指针 指向A的Random指向的那个节点的复制版
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			//node的Next复制节点的Random指向 node的Random指针指向的那个节点的Next(复制节点)
			node.Next.Random = node.Random.Next
		}
	}
	//拆分链表
	//原始结构
	// A -> A` -> B ->B` ->C ->C`
	// 新链表A`->B`->C`
	newHead := head.Next
	for node := head; node != nil; node = node.Next {
		newNode := node.Next
		node.Next = newNode.Next
		if newNode.Next != nil {
			newNode.Next = newNode.Next.Next
		}
	}

	return newHead
}