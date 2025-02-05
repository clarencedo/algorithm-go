package leetcode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

// INFO: 寻找下一个更大的元素，经典的单调栈问题
func nextLargerNodes(head *ListNode) []int {
	var res []int
	var stack [][]int
	cur := head
	idx := -1

	for cur != nil {
		idx++
		res = append(res, 0)
		for len(stack) > 0 && stack[len(stack)-1][0] < cur.Val {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top[1]] = cur.Val
		}
		stack = append(stack, []int{cur.Val, idx})
		cur = cur.Next
	}

	return res
}