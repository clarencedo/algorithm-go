package leetcode

func twoSum167(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

func twoSumII(numbers []int, target int) []int {
	right := len(numbers) - 1
	for left := 0; left < right; left++ {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		for sum > target {
			right--
		}
	}
	return []int{-1, -1}
}