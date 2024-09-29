package leetcode

import (
	"clarencedu/algorithm-go/structure"
	"sort"
	"fmt"
)
type TreeNode = structure.TreeNode

// type Pair = structure.Pair

// 解法一： BFS， 每个节点记录一个 行列坐标 row,column
// root 为0,0 ，左字节点为1,-1，右字节点为1,1
func verticalOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	columnTable := make(map[int][]int)
	queue := []structure.Pair[*TreeNode, int]{{Key: root, Value: 0}}
	column := 0

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		node := p.Key
		column = p.Value
		if node != nil {
			if _, exists := columnTable[column]; !exists {
				columnTable[column] = []int{}
			}
			columnTable[column] = append(columnTable[column], node.Val)

			queue = append(queue, structure.Pair[*structure.TreeNode, int]{Key: node.Left, Value: column - 1})
			queue = append(queue, structure.Pair[*structure.TreeNode, int]{Key: node.Right, Value: column + 1})
		}
	}
	sortedKeys := make([]int, 0, len(columnTable))
	for k := range columnTable{}
		sortedKeys = append(sortedKeys, k)
	}
	for _ ,k := range sortedKeys{
	res = append(res, columnTable[k])
	}

return res
}