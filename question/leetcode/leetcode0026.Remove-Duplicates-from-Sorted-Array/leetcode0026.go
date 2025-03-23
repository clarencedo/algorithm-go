package leetcode

func removeDuplicates(nums []int) int {
	//0索引 跳过
	slow, fast := 1, 1
	for fast < len(nums) {
		// 与前一个元素不相等
		if nums[fast] != nums[slow-1] {
			// 将不相等的元素移到前面
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	// 结束的时候slow指向的是最后一个不重复的元素
	return slow
}
