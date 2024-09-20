package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
	// return bfs(root)
}

func dfs(root *TreeNode, curVal int) int {
	if root == nil {
		return 0
	}
	curVal = curVal*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	return dfs(root.Left, curVal) + dfs(root.Right, curVal)
}

func bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left == nil && node.Right == nil {
			res += node.Val
		}
		if node.Left != nil {
			node.Left.Val += node.Val * 10
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			node.Right.Val += node.Val * 10
			queue = append(queue, node.Right)
		}
	}
	return res
}