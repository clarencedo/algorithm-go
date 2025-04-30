package leetcode

// 连续多个target，返回第一个和最后一个的下标
func searchRange(nums []int, target int) []int {
	leftIndex := leftBound(nums, target)
	rightIndex := rightBound(nums, target)
	return []int{leftIndex, rightIndex}
}

// 找target左边的边界
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
	//循环结束时，left指向第一个大于target的元素，right指向最后一个小雨target的元素
	//检查left是否越界，可能所有元素都小于target
	// nums[2,4,6] ,target =5
	// 循环结束, left =2 ,right = 1
	if left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

// 找target右边的边界
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
	//循环结束时，left指向第一个大于target的元素，right指向最后一个小于等于target的元素
	//检查right是否越界，可能所有元素都大于target
	// nums[2,4,6] ,target =5
	// 循环结束, left = 0 ,right = -1
	if right >= 0 && nums[right] == target {
		return right
	}
	return -1
}