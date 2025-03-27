package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

func rob337(root *TreeNode) int {
	robThis, notRobThis := dfs(root)
	return max(robThis, notRobThis)
}

func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftRob, leftNotRob := dfs(root.Left)
	rightRob, rightNotRob := dfs(root.Right)

	robThis := root.Val + leftNotRob + rightNotRob
	notRobThis := max(leftRob, leftNotRob) + max(rightRob, rightNotRob)

	return robThis, notRobThis
}