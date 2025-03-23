package leetcode

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	capacity := 0
	for left < right {
		curr := min(height[left], height[right]) * (right - left)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		if curr > capacity {
			capacity = curr
		}
	}

	return capacity
}