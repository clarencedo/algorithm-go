package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

// bfs，从右往左入队列， 这样队列清空之后最后一个节点就是最左边的节点
// 注意避免修改root，要用node来存储当前节点的状态，而不修改root
// 同理：找最右边的节点，从左往右入队列
func findBottomLeftValue(root *TreeNode) int {
	node := root
	q := []*TreeNode{node}

	for len(q) > 0 {
		node, q = q[0], q[1:]
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
	}
	return node.Val
}
