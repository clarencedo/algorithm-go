package leetcode

import "sort"

func canBeEqual(target []int, arr []int) bool {
	mp := map[int]int{}
	for i, v := range target {
		mp[v]++
		mp[arr[i]]--
	}

	for _, char := range mp {
		if char != 0 {
			return false
		}
	}
	return true
}

func canBeEqualII(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)
	for i := range target {
		if target[i] != arr[i] {
			return false
		}
	}
	return true
}
