package leetcode

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)

	//第一行和第一列的路径数都是1,只有一种走法
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// 空间复杂度O(n)
// 一纬数组滚动更新
func uniquePathsII(m int, n int) int {
	dp := make([]int, n)

	// 第一行的路径数都是1
	for j := range dp {
		dp[j] = 1
	}

	for i := 1; i < m; i++ {
		// 第一列始终为1
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1] // 当前位置路径数 = 上方路径数(dp[j]) + 左侧路径数(dp[j-1])
		}
	}

	return dp[n-1]
}