package leetcode

// 前缀和+ hashmap
// 把问题转换成 找前缀和相同的两个下标，他们之间的01个数相等
// 技巧：把0->-1，这样问题就变成了 找前缀和相同的两个下标，他们之间的1的个数相等
func findMaxLength(nums []int) int {
	//key为前缀和,val为前缀和出现的下标
	prefixIndex := make(map[int]int)
	prefixIndex[0] = -1

	maxLen := 0
	prefixSum := 0

	for i, num := range nums {
		if num == 0 {
			prefixSum += -1
		} else {
			prefixSum += 1
		}

		if prevIndex, ok := prefixIndex[prefixSum]; ok {
			maxLen = max(maxLen, i-prevIndex)
		} else {
			prefixIndex[prefixSum] = i
		}
	}

	return maxLen
}
