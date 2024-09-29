package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

// 解法一
func countUnivalSubtrees(root *TreeNode) int {
	count := 0
	dfs(root, &count)

	return count
}

func dfs(node *TreeNode, count *int) bool {
	if node == nil {
		return true
	}
	isUnival := true
	if node.Left != nil {
		if !dfs(node.Left, count) || node.Val != node.Left.Val {
			isUnival = false
		}
	}
	if node.Right != nil {
		if !dfs(node.Right, count) || node.Val != node.Right.Val {
			isUnival = false
		}
	}
	if isUnival {
		*count++
	}
	return isUnival
}

// 解法二
// 不检查每个节点是否有字节点，而是将空节点也堪称同值子树，但是空节点不计数
func countUnivalSubtreesII(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	dfsII(root, root.Val, &count)
	return count
}

func dfsII(node *TreeNode, val int, count *int) bool {
	if node == nil {
		return true
	}
	flagLeft := dfsII(node.Left, node.Val, count)
	falgRight := dfsII(node.Right, node.Val, count)
	if flagLeft && falgRight {
		*count++
		return node.Val == val
	} else {
		return false
	}
}