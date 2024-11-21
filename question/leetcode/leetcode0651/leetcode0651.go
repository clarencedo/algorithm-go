package leetcode

func maxA(n int) int {
	dp := make([]int, n+1) // dp[i] 表示最多 i 次按键后的最大 A 数量

	for i := 1; i <= n; i++ {
		// 1. 按下一个 'A'
		dp[i] = dp[i-1] + 1

		// 2. 尝试在第 j 次按键后切换到 Ctrl-A + Ctrl-C + 若干次 Ctrl-V
		for j := 1; j <= i-2; j++ {
			dp[i] = max(dp[i], dp[j]*(i-j-1))
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}