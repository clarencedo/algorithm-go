package leetcode

import "sort"

func twoSum(nums []int, target int) []int {
	// key: num, value: index
	table := make(map[int]int)

	for i, v := range nums {
		if p, ok := table[target-v]; ok {
			return []int{p, i}
		}
		table[v] = i
	}
	return nil
}

func twoSumII(nums []int, target int) []int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}