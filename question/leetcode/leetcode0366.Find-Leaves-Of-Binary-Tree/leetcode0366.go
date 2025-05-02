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

// 收集所有叶子节点
// 移除所有叶子节点
// 直到树为空
var res [][]int

func findLeavesII(root *TreeNode) [][]int {
	res = [][]int{}
	dfs(root)
	return res
}

func dfs(node *TreeNode) int {
	if node == nil {
		return -1
	}
	//深度从叶子节点往上计算，叶子节点深度为0
	depth := 1 + max(dfs(node.Left), dfs(node.Right))
	if depth == len(res) {
		res = append(res, []int{})
	}
	//将当前节点的值添加到对应深度的切片中
	res[depth] = append(res[depth], node.Val)
	return depth
}
