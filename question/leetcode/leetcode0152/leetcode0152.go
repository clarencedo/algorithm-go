package leetcode

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxProd, minProd, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		// 当前值可能使之前的最小乘积变为最大乘积，因此需要暂存当前的最大乘积
		tempMax := maxProd
		//有负号的原因，最大乘积 可能是  最大乘积* 当前值， 也可能是最小乘积*当前值， 也可以就是当前值
		maxProd = max(max(tempMax*nums[i], minProd*nums[i]), nums[i])
		//同理，最小乘积可能是 当前位置的最大乘积 * 当前值， 也可能是最小乘积*当前值， 也可能是当前值
		minProd = min(min(tempMax*nums[i], minProd*nums[i]), nums[i])
		result = max(result, maxProd)
	}

	return result
}