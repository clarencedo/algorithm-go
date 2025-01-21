package leetcode

import (
	"clarencedu/algorithm-go/structure"
	"math"
)

type TreeNode = structure.TreeNode

// 方法一： 递归 O(n^2)
func largestBSTSubtree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if isBST(root, nil, nil) {
		return count(root)
	}
	return max(largestBSTSubtree(root.Left), largestBSTSubtree(root.Right))
}

// 判断是否是二叉搜索树
func isBST(root *TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= *min {
		return false
	}
	if max != nil && root.Val >= *max {
		return false
	}
	return isBST(root.Left, min, &root.Val) && isBST(root.Right, &root.Val, max)
}

// 计算节点个数
func count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + count(root.Left) + count(root.Right)
}

// 解法二： 后续遍历 + 动态规划 O(n)
func largestBSTSubtreeII(root *TreeNode) int {
	maxSize := 0

	var postorder func(node *TreeNode) (bool, int, int, int)
	postorder = func(node *TreeNode) (bool, int, int, int) {
		if node == nil {
			return true, 0, math.MaxInt64, math.MinInt64
		}

		leftIsBST, leftSize, leftMin, leftMax := postorder(node.Left)

		rightIsBST, rightSize, rightMin, rightMax := postorder(node.Right)

		// 判断当前节点是否是BST
		if leftIsBST && rightIsBST && leftMax < node.Val && rightMin > node.Val {
			size := leftSize + rightSize + 1
			if size > maxSize {
				maxSize = size
			}
			return true, size, min(node.Val, leftMin), max(node.Val, rightMax)
		}

		return false, max(leftSize, rightSize), 0, 0
	}

	postorder(root)
	return maxSize
}