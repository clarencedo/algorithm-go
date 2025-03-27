package leetcode

func maxProfit309(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// dp[i][0]: 受伤持有股票的最大利润
	// dp[i][1]: 不持有股票，且处于冷冻期的最大利润
	// dp[i][2]: 不持有股票，且不处于冷冻期的最大利润
	// 状态转移方程：
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		//前一天持有股票，今天不动 或者 前一天不持有股票，今天买入
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		//前一天持有股票，今天卖出
		dp[i][1] = dp[i-1][0] + prices[i]
		//前一天不持有股票，且不处于冷冻期 或者 前一天处于冷冻期
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[n-1][1], dp[n-1][2])
}