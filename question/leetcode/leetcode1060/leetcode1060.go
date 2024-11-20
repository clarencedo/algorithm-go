package leetcode

func missingElement(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid]-nums[0]-mid < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left]-nums[0]-left < k {
		return nums[left] + k - (nums[left] - nums[0] - left)
	}
	return nums[0] + left + k

}