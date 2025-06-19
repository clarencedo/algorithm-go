package leetcode

import "sort"

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2

	// 找到第一个比右边数字小的数字
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// 如果不存在这样的数字，说明当前排列已经是最大的字典序，直接将数组逆序即可
	if i < 0 {
		sort.Ints(nums)
		return
	}

	// 从右向左再次扫描数组，找到第一个比 i 处数字大的数字
	j := n - 1
	for j >= 0 && nums[j] <= nums[i] {
		j--
	}

	// 交换 i 和 j 处的数字
	nums[i], nums[j] = nums[j], nums[i]

	// 将 i 后面的数字逆序排列
	// 灵魂就是这个反转
	reverse(nums[i+1:])
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}
