package leetcode

import "math"

func jump(nums []int) int {
	dp := make([]int, len(nums))
	for i := range nums {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(nums)-1]
}

// 贪心
func jumpII(nums []int) int {
	jump, currentEnd, farthest := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		// if i+nums[i] > farthest {
		// 	farthest = i + nums[i]
		// }
		//
		farthest = max(farthest, i+nums[i])
		//当前跳跃的边界
		if i == currentEnd {
			jump++
			currentEnd = farthest
		}
	}
	return jump
}