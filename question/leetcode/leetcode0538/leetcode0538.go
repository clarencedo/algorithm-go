package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

// 根据定义，dfs按照 右-根-左的顺序遍历二叉树，累加节点值，更新节点值
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}

	dfs(root)
	return root
}
