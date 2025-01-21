package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var head, pre *TreeNode
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != nil {
			pre.Right = node
		} else {
			head = node
		}
		node.Left = pre
		pre = node
		dfs(node.Right)
	}
	dfs(root)
	head.Left = pre
	pre.Right = head
	return head
}
