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

func postorder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	postorder(node.Left, res)
	postorder(node.Right, res)
	*res = append(*res, node.Val)
}

// 解法一 递归
func preorderTraversalII(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		res = append(res, root.Val)
		tmp := preorderTraversal(root.Left)
		for _, t := range tmp {
			res = append(res, t)
		}
		tmp = preorderTraversal(root.Right)
		for _, t := range tmp {
			res = append(res, t)
		}
	}
	return res
}

// 解法二 递归
func preorderTraversal1(root *TreeNode) []int {
	var result []int
	preorderII(root, &result)
	return result
}

func preorderII(root *TreeNode, output *[]int) {
	if root != nil {
		*output = append(*output, root.Val)
		preorderII(root.Left, output)
		preorderII(root.Right, output)
	}
}

// 解法三 非递归，用栈模拟递归过程
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack, res := []*TreeNode{}, []int{}
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			res = append(res, node.Val)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}
