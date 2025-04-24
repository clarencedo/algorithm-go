package leetcode

import (
	"clarencedu/algorithm-go/structure"
	"math/rand"
)

type ListNode = structure.ListNode

type Solution []int

func Constructor(head *ListNode) (s Solution) {
	for node := head; node != nil; node = node.Next {
		s = append(s, node.Val)
	}
	return s
}

func (s Solution) GetRandom() int {
	return s[rand.Intn(len(s))]
}