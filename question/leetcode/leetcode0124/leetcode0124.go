package leetcode

import (
	"clarencedu/algorithm-go/structure"
	"math"
)

type TreeNode = structure.TreeNode

var maxSum int

func maxPathSum(root *TreeNode) int {
	maxSum = math.MinInt32
	dfs(root)
	return maxSum
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(0, dfs(root.Left))
	right := max(0, dfs(root.Right))
	maxSum = max(maxSum, left+root.Val+right)
	return root.Val + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func maxPathSum(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	max := math.MinInt32
// 	getPathSum(root, &max)
// 	return max
// }
//
// func getPathSum(root *TreeNode, maxSum *int) int {
// 	if root == nil {
// 		return math.MinInt32
// 	}
// 	left := getPathSum(root.Left, maxSum)
// 	right := getPathSum(root.Right, maxSum)
//
// 	currMax := max(max(left+root.Val, right+root.Val), root.Val)
// 	*maxSum = max(*maxSum, max(currMax, left+right+root.Val))
// 	return currMax
// }
//
// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }