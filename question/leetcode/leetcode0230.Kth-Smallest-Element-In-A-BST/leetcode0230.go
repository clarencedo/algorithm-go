package leetcode

import (
	"clarencedu/algorithm-go/structure"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structure.TreeNode

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	var res []int

	dfs230(root, &res)
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res[k-1]
}

func dfs230(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	if root.Left != nil {
		dfs230(root.Left, res)
	}
	if root.Right != nil {
		dfs230(root.Right, res)
	}
}

// 定义中说了是二叉搜索树，所以可以用中序遍历，然后取第k个值,不需要再去循环排序了
func kthSmallestII(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	var res []int

	inorder(root, &res)
	return res[k-1]
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}