package question

func searchRange(nums []int, target int) []int {
	leftIndex := leftBound(nums, target)
	rightIndex := rightBound(nums, target)
	return []int{leftIndex, rightIndex}
}

func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 缩减右边界,继续往左边找,锁定左侧边界
			right = mid - 1
		}
	}
	if left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 缩减左边界,继续往右边找,锁定右侧边界
			left = mid + 1
		}
	}
	if right >= 0 && nums[right] == target {
		return right
	}
	return -1
}