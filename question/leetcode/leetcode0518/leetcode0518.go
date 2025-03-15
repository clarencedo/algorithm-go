package leetcode

// 零钱兑换II
// 返回凑成总金额为amount的硬币组合数
func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}

	//凑成i有多少种方式
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for x := coin; x <= amount; x++ {
			//在选择x金额之前，x-coin金额有dp[x-coin]种方式
			dp[x] += dp[x-coin]
		}
	}
	return dp[amount]
}