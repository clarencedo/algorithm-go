package leetcode

// 时间复杂度：O(m*n)
// 空间复杂度：O(m*n)
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	//dp[i][j]表示以matrix[i][j]为右下角的正方形的最大边长
	//每次更新最大边长,遍历到最右下角的时候，最大边长即为所求
	//状态转移方程：dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
	//找左边 上边 左上边 三个位置的最小值+1
	dp := make([][]int, m)
	maxSide := 0

	//初始化dp数组
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	//从左上角开始遍历
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				maxSide = max(maxSide, dp[i][j])
			}
		}
	}

	return maxSide * maxSide
}