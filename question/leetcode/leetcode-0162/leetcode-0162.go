package leetcode

import "math"

// 转换为查找mid>mid+1的元素，找了之后
// mid不一定是峰值，有可能它正处于下降的区间，mid左边的值大于mid
// 此时移动右边指针，直至left = right
func findPeakElement(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		mid := (right-left)>>1 + left
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func findPeakElementII(nums []int) int {
	n := len(nums)
	//因为nums[-1] = nums[n] = -∞,所以规避第一个元素比第二个元素大，最后一个元素比倒数第二个元素小的情况。
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)>>1 + left
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0

}