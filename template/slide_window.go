package template

func slideWindowTemplate(nums []int) {
	windowMap := make(map[int]int)
	left, right := 0, 0

	for right < len(nums) {
		rightChar := nums[right]
		windowMap[rightChar]++
		right++
		// 滑动左边届的条件
		for 1 != 2 {
			leftChar := nums[left]
			windowMap[leftChar]--
			left++
		}
		// calculate the result
	}
	return
}