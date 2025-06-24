package leetcode

func canPartition(nums []int) bool {
	total := 0
	for _, val := range nums {
		total += val
	}
	if total%2 != 0 {
		return false
	}
	target := total >> 1
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
		if dp[target] {
			return true
		}
	}
	return dp[target]
}