package leetcode

// slide window
func findMaxConsecutiveOnes(nums []int) int {
	window := make(map[int]int)
	n := len(nums)
	left, right := 0, 0
	maxLen := 0
	for right < n {
		rightVal := nums[right]
		window[rightVal]++
		right++
		for window[0] > 1 {
			leftVal := nums[left]
			window[leftVal]--
			left++
		}
		if right-left > maxLen {
			maxLen = right - left
		}
	}
	return maxLen
}

// convert int into string , spilit with '0',calculate the longest array with consecutive '1'