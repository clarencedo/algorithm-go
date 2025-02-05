package leetcode

// 解法一： 双指针
// 时间复杂度 O(n^2)
// 空间复杂度 O(1)
func trap(height []int) int {
	sum := 0
	for i := 1; i < len(height)-1; i++ {
		maxLeft, maxRight, min := 0, 0, 0
		// 向左找最高的左边界
		for j := i - 1; j >= 0; j-- {
			if height[j] > maxLeft {
				maxLeft = height[j]
			}
		}
		// 向右找最高的右界
		for j := i + 1; j < len(height); j++ {
			if height[j] > maxRight {
				maxRight = height[j]
			}
		}
		if maxLeft < maxRight {
			min = maxLeft
		} else {
			min = maxRight
		}
		// 只有min>当前柱子的高度，才能计算每个柱子上方可以储水的量
		if min > height[i] {
			sum += (min - height[i])
		}
	}
	return sum
}

// 解法二： 双指针优化，一次遍历
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func trapII(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return
}

// 解法二： 动态规划
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func trapIII(height []int) (ans int) {
	n := len(height)
	if n == 0 {
		return
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	for i, h := range height {
		ans += min(leftMax[i], rightMax[i]) - h
	}
	return
}

// 解法三： 单调栈
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func trapIV(height []int) (ans int) {
	stack := []int{}
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			ans += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return
}
