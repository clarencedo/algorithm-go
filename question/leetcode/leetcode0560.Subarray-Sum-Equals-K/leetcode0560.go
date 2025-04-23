package leetcode

func subarraySum(nums []int, k int) int {
	preSum, count := 0, 0
	// key: 前缀和, value: 出现次数
	m := make(map[int]int)
	m[0] = 1

	for _, num := range nums {
		preSum += num
		//如果preSum[j]- preSum[i] = k
		//那么[i+1,...,j]这一段子数组的和为k
		if _, ok := m[preSum-k]; ok {
			count += m[preSum-k]
		}
		m[preSum]++
	}

	return count
}
