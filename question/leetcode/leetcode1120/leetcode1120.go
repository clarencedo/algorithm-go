package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

func maximumAverageSubtree(root *TreeNode) float64 {
	var max float64
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		lsum, lcnt := dfs(node.Left)
		rsum, rcnt := dfs(node.Right)
		sum := lsum + rsum + node.Val // 以node为根的子树的节点值之和
		cnt := lcnt + rcnt + 1        // 以node为根的子树的节点个数
		max = Max(max, float64(sum)/float64(cnt))
		return sum, cnt
	}
	dfs(root)
	return max
}

func Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}