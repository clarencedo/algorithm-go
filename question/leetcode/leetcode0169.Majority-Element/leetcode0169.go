package leetcode

func majorityElement(nums []int) int {
	//candidate：用于存储当前的候选多数元素
	//count：用于记录当前候选元素的票数
	count := 0
	candidate := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count += 1
		} else {
			count -= 1
		}
	}

	return candidate
}