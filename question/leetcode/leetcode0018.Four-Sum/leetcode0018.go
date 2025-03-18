package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int
	for i := 0; i < -n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i >= 0 && nums[i] == nums[i-1] {
			//去重
			continue
		}
		for j := i + 1; i < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				//去重
				continue
			}
			left, right := j+1, n-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}

	}
	return ans
}