package leetcode

// 全排列，返回所有可能的全排列
func permute(nums []int) [][]int {
	res := [][]int{}
	track := []int{}
	used := make([]bool, len(nums))
	backtrack(nums, track, used, &res)
	return res
}
func backtrack(nums []int, track []int, used []bool, res *[][]int) {
	if len(track) == len(nums) {
		*res = append(*res, append([]int{}, track...))
		return
	}

	for i := range nums {
		if used[i] {
			continue
		}
		track = append(track, nums[i])
		used[i] = true
		backtrack(nums, track, used, res)
		track = track[:len(track)-1]
		used[i] = false
	}
}

// 全排列II，返回不重复的全排列
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	track := []int{}
	used := make([]bool, len(nums))
	backtrackII(nums, track, used, &res)
	return res
}

func backtrackII(nums []int, track []int, used []bool, res *[][]int) {
	if len(track) == len(nums) {
		*res = append(*res, append([]int{}, track...))
		return
	}
	// used标记全局已使用的元素，duplicated标记当前 同层 已使用的元素
	duplicated := make(map[int]bool)
	for i := range nums {
		if _, ok := duplicated[nums[i]]; !ok && !used[i] {
			duplicated[nums[i]] = true
			used[i] = true
			track = append(track, nums[i])

			backtrackII(nums, track, used, res)
			used[i] = false
			track = track[:len(track)-1]
		}
	}
}