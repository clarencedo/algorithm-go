package leetcode

import "math"

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left, right := 0, 0
	ans := math.MaxInt32
	sum := 0
	for right < n {
		sum += nums[right]
		right++
		for sum >= target {
			ans = min(ans, right-left)
			sum -= nums[left]
			left++
		}
	}

	if ans == math.MaxInt32 {
		return 0
	}

	return ans
}
