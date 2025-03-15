package leetcode

import "sort"

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	track := []int{}
	backtrack(0, track, nums, &res)
	return res
}

// 因为已经知晓树族中的元素 互不相同，所以不需要used数组
func backtrack(index int, current []int, nums []int, res *[][]int) {
	tmp := append([]int{}, current...)
	*res = append(*res, tmp)

	for i := index; i < len(nums); i++ {
		current = append(current, nums[i])
		backtrack(i+1, current, nums, res)
		current = current[:len(current)-1]
	}
}

// 如果有重复元素,每一层判断以下是否重复即可
// 因为已经排序了，所以只需要判断当前元素是否和上一个元素相同即可
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	track := []int{}
	backtrackII(0, nums, track, &res)
	return res
}

func backtrackII(index int, nums, track []int, res *[][]int) {
	tmp := append([]int{}, track...)
	*res = append(*res, tmp)

	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		track = append(track, nums[i])
		backtrackII(i+1, nums, track, res)
		track = track[:len(track)-1]
	}
}