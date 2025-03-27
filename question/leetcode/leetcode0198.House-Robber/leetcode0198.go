package leetcode

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}

// 空间优化
func robII(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 0 {
		return 0
	}
	prePre, pre := nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		cur := max(pre, prePre+nums[i])
		prePre, pre = pre, cur
	}
	return pre
}

// 213题 与 198题的区别
// 环形街区，第一个房子和最后一个房子相邻
// 可以转化为两个单排街区，第一个房子偷或者不偷
// 第一个房子偷，最后一个房子不偷
// 第一个房子不偷，最后一个房子偷
func robIII(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 0 {
		return 0
	}
	return max(robIndex(nums, 0, n-2), robIndex(nums, 1, n-1))
}

func robIndex(nums []int, start, end int) int {
	n := end - start + 1
	if n == 1 {
		return nums[start]
	}
	if n == 0 {
		return 0
	}
	prePre, pre := nums[start], max(nums[start], nums[start+1])
	for i := start + 2; i <= end; i++ {
		cur := max(pre, prePre+nums[i])
		prePre, pre = pre, cur
	}
	return pre
}
