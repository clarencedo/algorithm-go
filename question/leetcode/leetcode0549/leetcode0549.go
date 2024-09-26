package leetcode

import (
	"clarencedu/algorithm-go/structure"
)

type TreeNode = structure.TreeNode

func longestConsecutive(root *TreeNode) int {
	maxLength := 0
	dfs(root, &maxLength)
	return maxLength
}

func dfs(node *TreeNode, maxLength *int) (int, int) {
	if node == nil {
		return 0, 0
	}

	// 获取左子树和右子树的连续序列长度
	leftInc, leftDec := dfs(node.Left, maxLength)
	rightInc, rightDec := dfs(node.Right, maxLength)

	inc, dec := 1, 1 // 初始化当前节点的向上和向下连续序列长度为 1

	// 检查左子节点
	if node.Left != nil {
		if node.Left.Val == node.Val+1 {
			inc = leftInc + 1 // 如果左子节点比当前节点大 1，更新向上长度
		} else if node.Left.Val == node.Val-1 {
			dec = leftDec + 1 // 如果左子节点比当前节点小 1，更新向下长度
		}
	}

	// 检查右子节点
	if node.Right != nil {
		if node.Right.Val == node.Val+1 {
			inc = max(inc, rightInc+1) // 更新向上长度
		} else if node.Right.Val == node.Val-1 {
			dec = max(dec, rightDec+1) // 更新向下长度
		}
	}

	// 更新全局最长长度，减去 1 是因为当前节点在两个序列中被计算了一次
	*maxLength = max(*maxLength, inc+dec-1)
	return inc, dec // 返回当前节点的向上和向下长度
}