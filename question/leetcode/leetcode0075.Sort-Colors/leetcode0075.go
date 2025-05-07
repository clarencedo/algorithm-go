package leetcode

// sortColors 荷兰国旗问题
// 三指针，left指向0的最右边界，right指向2的最左边界
// 遍历数组，如果遇到0则与left交换，left右移；如果遇到2则与right交换，right左移
// 时间复杂度O(n)，空间复杂度O(1)
func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	for i := 0; i <= right; i++ {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
			i--
		}
	}
}
