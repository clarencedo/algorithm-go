package leetcode

func longestConsecutive(nums []int) int {
	//哈希表记录数字是否存在
	numMap := map[int]bool{}
	for _, val := range nums {
		numMap[val] = true
	}
	ans := 0
	for num := range numMap {
		//当numMap[num-1]不存在，才认为当前是一个连续序列的起点
		if !numMap[num-1] {
			current := num
			currentStreak := 1
			for numMap[current+1] {
				current++
				currentStreak++
			}
			if ans < currentStreak {
				ans = currentStreak
			}
		}
	}
	return ans
}

// 解法三 暴力解法，时间复杂度 O(n^2)
func longestConsecutive2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numMap, length, tmp, lcs := map[int]bool{}, 0, 0, 0
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = true
	}
	for key := range numMap {
		if !numMap[key-1] && !numMap[key+1] {
			delete(numMap, key)
		}
	}
	if len(numMap) == 0 {
		return 1
	}
	for key := range numMap {
		if !numMap[key-1] && numMap[key+1] {
			length, tmp = 1, key+1
			for numMap[tmp] {
				length++
				tmp++
			}
			lcs = max(lcs, length)
		}
	}
	return max(lcs, length)
}
