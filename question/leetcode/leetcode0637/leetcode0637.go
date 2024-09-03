package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	res := []float64{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		var currentLevelVal float64 = 0.00
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevelVal += float64(node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, currentLevelVal/float64(levelSize))
	}

	return res
}