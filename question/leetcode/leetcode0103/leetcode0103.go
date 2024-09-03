package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}
	queue := []*TreeNode{root}
	fromLeft := true
	for len(queue) > 0 {
		levelSize := len(queue)
		levelVal := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if fromLeft {
				levelVal[i] = node.Val
			} else {
				levelVal[levelSize-1-i] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, levelVal)
		fromLeft = !fromLeft
	}
	return res

}