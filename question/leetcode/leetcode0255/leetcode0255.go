package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

func verifyPreorder(preorder []int) bool {
	stack := make([]int, 0)
	low := -1 << 31
	for _, v := range preorder {
		if v < low {
			return false
		}
		for len(stack) > 0 && v > stack[len(stack)-1] {
			low = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
	return true

}