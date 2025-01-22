package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

// 根据性质，只需要判断有效节点之前是否有空节点即可
// 用层序遍历，遇到空节点标记flag，后续节点不能有非空节点
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var queue []*TreeNode
	queue = append(queue, root)
	var flag bool

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			flag = true
		} else {
			if flag {
				return false
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	return true
}
