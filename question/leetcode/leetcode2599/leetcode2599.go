package leetcode

import "sort"

// 方法 ：贪心 + 逆序排序
// •	核心思想：当前缀和变为负数时，将 nums 中最小的负数移到数组末尾。
// •	实现方式：
// 1.	遍历 nums 计算前缀和，一旦前缀和小于 0，记录需要调整的负数。
// 2.	对所有记录的负数进行降序排序，然后按需移除，直到前缀和非负。
// 3.	时间复杂度：O(n log n)（排序）。
func makePrefSumNonNegative(nums []int) int {
	moveCount := 0
	negNums := make([]int, 0)
	prefix := 0

	for _, num := range nums {
		prefix += num
		if num < 0 {
			negNums = append(negNums, num)
		}

		if prefix < 0 {
			sort.Slice(negNums, func(i, j int) bool {
				return negNums[i] > negNums[j]
			})
			prefix = prefix - negNums[0]
			negNums = negNums[1:]
			moveCount++
		}

	}

	return moveCount
}