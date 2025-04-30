package leetcode

import "clarencedu/algorithm-go/structure"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structure.TreeNode

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val > root.Val {
		// 新值比当前节点大，新值成为新的根节点
		return &TreeNode{Val: val, Left: root}
	}

	// 否则，递归到右子树继续插入
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}