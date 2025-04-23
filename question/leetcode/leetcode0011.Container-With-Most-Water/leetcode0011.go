package leetcode

// 解法一：双指针
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

// 解法而：暴力解
func maxAreaII(height []int) int {
	capacity := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			water := min(height[i], height[j]) * (j - i)
			if water > capacity {
				capacity = water
			}

		}
	}
	return capacity
}