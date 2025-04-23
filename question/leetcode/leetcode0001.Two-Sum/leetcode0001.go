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

type Pair struct {
	Val int
	Idx int
}

func twoSumII(nums []int, target int) []int {
	n := len(nums)
	arr := make([]Pair, n)
	for i := 0; i < n; i++ {
		arr[i] = Pair{Val: nums[i], Idx: i}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Val < arr[j].Val
	})
	left, right := 0, n-1
	for left < right {
		sum := arr[left].Val + arr[right].Val
		if sum == target {
			return []int{arr[left].Idx, arr[right].Idx}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
