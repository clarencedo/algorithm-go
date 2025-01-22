package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

// 递归，左叶子节点之和
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			res += node.Left.Val
		}

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return res
}

// 递归，右叶子节点之和
func sumOfRightLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Right != nil && node.Right.Left == nil && node.Right.Right == nil {
			res += node.Right.Val
		}

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return res
}