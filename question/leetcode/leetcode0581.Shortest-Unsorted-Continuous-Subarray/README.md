# [581. Shortest Unsorted Continuous Subarray](https://leetcode.com/problems/shortest-unsorted-continuous-subarray/)

## 题目

Given an integer array `nums`, you need to find one **continuous subarray** that if you only sort this subarray in ascending order, then the whole array will be sorted in ascending order.

Return *the shortest such subarray and output its length*.

**Example 1:**

```
Input: nums = [2,6,4,8,10,9,15]
Output: 5
Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
```

**Example 2:**

```
Input: nums = [1,2,3,4]
Output: 0
```

**Example 3:**

```
Input: nums = [1]
Output: 0
```

**Constraints:**

- `1 <= nums.length <= 104`
- `105 <= nums[i] <= 105`

## 题目大意

给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。请你找出符合题意的 最短 子数组，并输出它的长度。

## 解题思路

- 本题求的是最短逆序区间。经过简单推理，可以知道，这个逆序区间一定由这个区间内的最小元素决定左边界，最大元素决定右边界。
- 先从左边找到第一个降序的元素，并记录最小的元素 min，再从右边往左找到最右边开始降序的元素，并记录最大的元素 max。最后需要还原最小元素和最大元素在原数组中正确的位置。以逆序区间左边界为例，如果区间外的一个元素比这个逆序区间内的最小元素还要小，说明它并不是左边界，因为这个小元素和逆序区间内的最小元素组合在一起，还是升序，并不是逆序。只有在左边区间外找到第一个大于逆序区间内最小元素，说明这里刚刚开始发生逆序，这才是最小逆序区间的左边界。同理，在逆序区间的右边找到第一个小于逆序区间内最大元素，说明这里刚刚发生逆序，这才是最小逆序区间的右边界。至此，最小逆序区间的左右边界都确定下来了，最短长度也就确定了下来。时间复杂度 O(n)，空间复杂度 O(1)。

## 代码

```go
package leetcode

import "math"

func findUnsortedSubarray(nums []int) int {
	n, left, right, minR, maxL, isSort := len(nums), -1, -1, math.MaxInt32, math.MinInt32, false
	// left
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			isSort = true
		}
		if isSort {
			minR = min(minR, nums[i])
		}
	}
	isSort = false
	// right
	for i := n - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			isSort = true
		}
		if isSort {
			maxL = max(maxL, nums[i])
		}
	}
	// minR
	for i := 0; i < n; i++ {
		if nums[i] > minR {
			left = i
			break
		}
	}
	// maxL
	for i := n - 1; i >= 0; i-- {
		if nums[i] < maxL {
			right = i
			break
		}
	}
	if left == -1 || right == -1 {
		return 0
	}
	return right - left + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```