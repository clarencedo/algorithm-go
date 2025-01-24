package leetcode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

type MyLinkedList struct {
	head *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{&ListNode{}, 0}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}

	cur := l.head

	for i := 0; i <= index; i++ {
		cur = cur.Next
	}

	return cur.Val
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.size {
		return
	}

	index = max(index, 0)
	l.size++
	preHead := l.head

	for i := 0; i < index; i++ {
		preHead = preHead.Next
	}

	toAdd := &ListNode{Val: val, Next: preHead.Next}
	preHead.Next = toAdd
}
func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}