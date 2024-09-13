package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

func lowestCommonAncestor235(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor235(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor235(root.Right, p, q)
	}
	return root
}