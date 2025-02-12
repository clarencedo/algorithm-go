package leetcode

import (
	"clarencedu/algorithm-go/structure"
	"math"
)

type TreeNode = structure.TreeNode

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt32
	getPathSum(root, &max)
	return max
}

func getPathSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return math.MinInt32
	}
	left := getPathSum(root.Left, maxSum)
	right := getPathSum(root.Right, maxSum)

	currMax := max(max(left+root.Val, right+root.Val), root.Val)
	*maxSum = max(*maxSum, max(currMax, left+right+root.Val))
	return currMax
}

var maxSum int

func maxPathSumII(root *TreeNode) int {
	maxSum := math.MinInt32
	dfs(root)
	return maxSum
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := max(dfs(root.Left), 0)
	right := max(dfs(root.Right), 0)
	currVal := root.Val + left + right
	maxSum = max(maxSum, currVal)
	return root.Val + max(left, right)
}