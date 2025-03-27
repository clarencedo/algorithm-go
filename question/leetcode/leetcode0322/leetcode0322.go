package leetcode

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) //凑成总金额为i所需的最少硬币个数
	for i := 0; i <= amount+1; i++ {
		//随便设置一个不可能的值
		dp[i] = amount + 1
	}

	dp[0] = 0
	//凑成amount的金额，从凑0开始迭代
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			//i-coin 为使用了当前面值的coin,当前金额减去当前硬币的值
			if i-coin < 0 {
				continue
			}
			//选中当前的硬币
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}
	//如果dp[amount] == amount+1(初始化的值)，说明没有凑出来,
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

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
