package leetcode

import (
	"clarencedu/algorithm-go/structure"
	"sort"
)

type TreeNode = structure.TreeNode

type NodeWithCol struct {
	node *TreeNode
	col  int
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	columnTable := make(map[int][]int)

	queue := []NodeWithCol{{root, 0}}

	minCol, maxCol := 0, 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		node, col := current.node, current.col

		columnTable[col] = append(columnTable[col], node.Val)

		if col < minCol {
			minCol = col
		}
		if col > maxCol {
			maxCol = col
		}

		if node.Left != nil {
			queue = append(queue, NodeWithCol{node.Left, col - 1})
		}
		if node.Right != nil {
			queue = append(queue, NodeWithCol{node.Right, col + 1})
		}

	}
	res := [][]int{}
	for col := minCol; col <= maxCol; col++ {
		res = append(res, columnTable[col])
	}

	return res
}

// BFS， 每个节点记录一个 行列坐标 row,column
// root 为0,0 ，左字节点为1,-1，右字节点为1,1
func verticalOrderII(root *TreeNode) [][]int {
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

	// sort.Ints(columnTable)
	sortedKeys := make([]int, 0, len(columnTable))
	for k := range columnTable {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Ints(sortedKeys)
	for _, k := range sortedKeys {
		res = append(res, columnTable[k])
	}
	return res
}
