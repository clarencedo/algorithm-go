package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

func preorderTraversal(root *TreeNode) []int {
	var res []int
	preorder(root, &res)
	return res
}

func preorder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	preorder(node.Left, res)
	preorder(node.Right, res)
}

func inorder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, res)
	*res = append(*res, node.Val)
	inorder(node.Right, res)
}