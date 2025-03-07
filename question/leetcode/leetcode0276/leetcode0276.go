package leetcode

func numWays(n int, k int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return k
	}
	if n == 2 {
		return k * k
	}
	// dp[i] means the number of ways to paint the first i posts
	dp := make([]int, n+1)
	dp[1] = k
	dp[2] = k * k

	// same[i]表示第i个柱子和第i-1个柱子颜色相同的涂色方案数
	// diff[i]表示第i个柱子和第i-1个柱子颜色不同的涂色方案数
	// dp[i] = same[i] + diff[i]
	// same[i] = diff[i-1] , 因为和i-1颜色相同，所以same[i]只有一种选择,就是i-1的颜色
	// diff[i] = (k-1) * (same[i-1] + diff[i-1]) =  (k-1) * dp[i-1] 因为和i-1颜色不同，所以diff[i]有k-1种选择，然后乘以dp[i-1]的方案数
	// 由此可以推导出diff[i] = (k-1) * dp[i-1] -> diff[i-1] = (k-1) * dp[i-2]
	// 所以由 dp[i] = diff[i-1] + (k-1) * dp[i-1]
	// 继续推出 dp[i] = (k-1) * dp[i-2] + (k-1) * dp[i-1]
	// 所以最后状态转移方程为 dp[i] = (k-1) * (dp[i-1] + dp[i-2])
	for i := 3; i <= n; i++ {
		dp[i] = (k - 1) * (dp[i-1] + dp[i-2])
	}

	return dp[n]
}

// 可以优化空间复杂度为O(1)
func numWaysII(n int, k int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return k
	}
	if n == 2 {
		return k * k
	}
	prev2 := k
	prev1 := k * k

	for i := 3; i <= n; i++ {
		cur := (k - 1) * (prev1 + prev2)
		prev2 = prev1
		prev1 = cur
	}

	return prev1
}
