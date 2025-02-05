package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

// INFO: BSF
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxWidth := 0
	queue := []*TreeNode{root}
	// 用于存储每个节点的位置
	levelQueue := []int{0}

	for len(queue) > 0 {
		size := len(queue)
		minIndex := levelQueue[0]
		maxIndex := levelQueue[size-1]
		maxWidth = max(maxWidth, maxIndex-minIndex+1)

		for i := 0; i < size; i++ {

			node := queue[i]
			index := levelQueue[i] - minIndex

			// 数组存储的方式，左子节点的索引是2*当前索引+1, 右子节点的索引是2*当前索引+2
			if node.Left != nil {
				queue = append(queue, node.Left)
				levelQueue = append(levelQueue, 2*index+1)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				levelQueue = append(levelQueue, 2*index+2)
			}
		}

		// 当前层已经处理过，从队列中移除
		queue = queue[size:]
		// 位置也移除
		levelQueue = levelQueue[size:]
	}
	return maxWidth
}
