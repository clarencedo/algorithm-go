package leetcode

import (
	"clarencedu/algorithm-go/structure"
	"math"
)

type TreeNode = structure.TreeNode

// 解法一： 按照定义比大小, O(n)
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isBST(root, math.MinInt64, math.MaxInt64)
}
func dfs(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}

	return isBST(node.Left, min, node.Val) && isBST(node.Right, node.Val, max)
}

// 解法二： 中序遍历,如果结果有逆序就不是BST, O(n) + O(n) = 0(n)
func isValidBSTII(root *TreeNode) bool {
	var res []int
	inorder(root, &res)
	for i := 1; i < len(res); i++ {
		if res[i-1] >= res[i] {
			return false
		}
	}
	return true
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}
