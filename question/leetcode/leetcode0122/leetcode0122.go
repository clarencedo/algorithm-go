package main

// 贪心
func maxProfit(prices []int) int {
	n := len(prices)
	profit := 0
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		} else {
			//do nothing
			continue
		}
	}
	return profit
}

// 动态规划
func maxProfitII(prices []int) int {
	n := len(prices)
	//定义状态 dp[i][0]表示第 i 天交易完后手里没有股票的最大利润，
	//dp[i][1] 表示第 i 天交易完后手里持有一支股票的最大利润
	dp := make([][2]int, n)
	//第0天没有股票dp[0][0] = 0, 有股票dp[0][0] = -prices[0]
	dp[0][0], dp[0][1] = 0, -prices[0]
	for i := 1; i < n; i++ {
		// 当天手里没有股票代表 前一天手里没有股票dp[i-1][0] ,或者 前一天有股票今天卖出 获得prices[i]收益
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 当天手里有股票代表 前一天手里已经有股票dp[i-1][1], 或者 前一天没有股票今天再买入股票dp[i-1][0] - prices[i]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func maxProfitIII(prices []int) int {
	n := len(prices)
	// dp0: 手上不持有股票的最大利润
	// dp1: 手上持有股票的最大利润
	// dp0 = max(dp0, dp1+prices[i])
	// dp1 = max(dp1, dp0-prices[i])
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
	}
	return dp0
}