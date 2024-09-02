package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, curVal int) int {
	if root == nil {
		return 0
	}
	curVal = curVal*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	return dfs(root.Left, curVal) + dfs(root.Right, curVal)
}