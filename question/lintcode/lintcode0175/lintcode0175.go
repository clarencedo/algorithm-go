package lintcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

func InvertBinaryTree(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	InvertBinaryTree(root.Left)
	InvertBinaryTree(root.Right)
}