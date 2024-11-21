package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

func boundaryOfBinaryTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	if !isLeaf(root) {
		res = append(res, root.Val)
	}
	leftBoundary(root.Left, &res)
	leaves(root, &res)
	rightBoundary(root.Right, &res)
	return res
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

// 递归遍历叶子节点
func leaves(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	if isLeaf(node) {
		*res = append(*res, node.Val)
		return
	}
	leaves(node.Left, res)
	leaves(node.Right, res)
}

func leftBoundary(node *TreeNode, res *[]int) {
	if node == nil || isLeaf(node) {
		return
	}
	*res = append(*res, node.Val)
	if node.Left != nil {
		leftBoundary(node.Left, res)
	} else {
		leftBoundary(node.Right, res)
	}
}

func rightBoundary(node *TreeNode, res *[]int) {
	if node == nil || isLeaf(node) {
		return
	}
	if node.Right != nil {
		rightBoundary(node.Right, res)
	} else {
		rightBoundary(node.Left, res)
	}
	*res = append(*res, node.Val)
}