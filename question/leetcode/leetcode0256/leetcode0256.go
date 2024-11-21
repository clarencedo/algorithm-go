package leetcode

import "math"

func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}
	n := len(costs)
	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, 3)
	}

	for i := 0; i < 3; i++ {
		dp[0][i] = costs[0][i]
	}

	for i := 1; i < n; i++ {
		for j := 0; j < 3; j++ {
			minCost := math.MaxInt32
			for pre := 0; pre < 3; pre++ {
				if pre != j {
					minCost = min(minCost, dp[i-1][pre]+costs[i][j])
				}
			}
			dp[i][j] = minCost
		}
	}

	lastroomMincost := math.MaxInt32
	for i := 0; i < 3; i++ {
		lastroomMincost = min(lastroomMincost, dp[n-1][i])
	}
	return lastroomMincost
}