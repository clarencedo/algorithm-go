const maxSubArrayII = (nums: Array<number>): number => {
	if (nums.length === 0) {
		return 0
	}

	let dp = new Array(nums.length)
	dp[0] = nums[0]
	let max = dp[0]
	for (let i = 1; i < nums.length; i++) {
		dp[i] = Math.max(dp[i - 1] + nums[i], nums[i])
		if (dp[i] > max) {
			max = dp[i]
		}
	}

	return max
}

