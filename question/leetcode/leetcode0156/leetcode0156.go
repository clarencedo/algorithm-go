package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}
	left := root.Left
	right := root.Right
	newRoot := upsideDownBinaryTree(left)
	left.Right = root
	left.Left = right
	root.Left = nil
	root.Right = nil

	return newRoot
}