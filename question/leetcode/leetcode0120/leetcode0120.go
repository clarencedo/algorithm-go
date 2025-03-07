package leetcode

// 从最后一行倒推,自底向上
// dp[j] 表示当前行j位置的最小路径和
// dp[j] 取决于 下一行的dp[i]和dp[i+1]
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)
	copy(dp, triangle[n-1])

	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
}