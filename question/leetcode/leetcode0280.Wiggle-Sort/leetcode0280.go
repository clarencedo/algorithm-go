package leetcode

import "sort"

// 解法一
// 倒序排列，交换奇数位和偶数位
func wiggleSort(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	n := len(nums)
	// 交换奇数位和偶数位
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
	return
}

// 解法二
// 贪心
func wiggleSortII(nums []int) {

}
func wiggleSortIII(nums []int) {
	for i := 1; i < len(nums); i++ {
		// 奇数位，nums[i] < nums[i-1]
		// 偶数位，nums[i] > nums[i-1]
		// 交换
		if (i%2 == 1 && nums[i] < nums[i-1]) || (i%2 == 0 && nums[i] > nums[i-1]) {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
}