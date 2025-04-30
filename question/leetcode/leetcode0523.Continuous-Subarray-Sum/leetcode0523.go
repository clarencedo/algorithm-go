package leetcode

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}

	// 记录前缀和的下标, key为前缀和，value为下标
	m := make(map[int]int)
	// 前缀和位0时，下标在-1
	m[0] = -1
	sum := 0
	//如果子数字[i+1,j]满足条件那么
	//sum[j] - sum[i] = n * k
	//只要两个前缀和取模之后的值相等，说明中间这段子数组的和是k的倍数
	for i, num := range nums {
		sum += num
		if k != 0 {
			sum %= k
		}
		if prevIndex, ok := m[sum]; ok {
			if i-prevIndex >= 2 {
				return true
			}
		} else {
			m[sum] = i
		}
	}
	return false
}