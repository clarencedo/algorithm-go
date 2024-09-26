package leetcode

import (
	"clarencedu/algorithm-go/structure"
	"math"
)

type TreeNode = structure.TreeNode

func longestConsecutive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return dfs(root, -1, math.MaxInt)
}

func dfs(node *TreeNode, parentVal, length int) int {
	if node == nil {
		return length
	}

	curLength := length
	if parentVal != math.MaxInt && node.Val == parentVal+1 {
		curLength++
	} else {
		curLength = 1
	}
	leftLen := dfs(node.Left, node.Val, curLength)
	rightLen := dfs(node.Right, node.Val, curLength)

	return max(curLength, max(leftLen, rightLen))
}