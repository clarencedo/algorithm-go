package leetcode

// 解法一：贪心,原地修改,nums[i]表示以i结尾的最大子数组和
func maxSubArray(nums []int) int {
	max := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

// 解法二：动态规划
func dp(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSum := dp[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > maxSum {
			maxSum = dp[i]
		}
	}

	return maxSum
}
