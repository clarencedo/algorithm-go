package leetcode

import (
	"clarencedu/algorithm-go/structure"
)

type TreeNode = structure.TreeNode

// 方法一： 正常层序遍历，然后将结果反转
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int

	queue := []*TreeNode{root}

	// 另一种写法
	// for len(queue) > 0 {
	// 	var level []int
	// 	var next []*TreeNode
	//
	// 	for _, node := range queue {
	// 		level = append(level, node.Val)
	// 		if node.Left != nil {
	// 			next = append(next, node.Left)
	// 		}
	// 		if node.Right != nil {
	// 			next = append(next, node.Right)
	// 		}
	// 	}
	//
	// 	res = append(res, level)
	// 	queue = next
	// }

	for len(queue) > 0 {
		size := len(queue)
		var currentNodes []int
		for i := 0; i < size; i++ {
			currentNode := queue[0]
			queue = queue[1:]
			currentNodes = append(currentNodes, currentNode.Val)
			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
		res = append(res, currentNodes)
	}

	return res
}