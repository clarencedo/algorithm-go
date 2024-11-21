package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

func findLeaves(root *TreeNode) [][]int {
	var res [][]int
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		depth := 1 + Max(dfs(node.Left), dfs(node.Right))
		if depth == len(res) {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], node.Val)
		return depth
	}
	dfs(root)
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}