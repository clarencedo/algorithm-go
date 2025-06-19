package leetcode

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	tmp := 1

	// 初始化ans数组
	for i := range ans {
		ans[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		tmp *= nums[i+1]
		ans[i] *= tmp
	}

	return ans
}
