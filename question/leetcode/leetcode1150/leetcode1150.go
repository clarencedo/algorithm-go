package leetcode

// 遍历
func isMajorityElement(nums []int, target int) bool {
	count := 0
	for _, num := range nums {
		if num == target {
			count++
		}
	}

	return count > len(nums)/2
}

// 遍历直到目标数字
func isMajorityElementII(nums []int, target int) bool {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			count++
		} else if nums[i] > target {
			break
		}
	}

	return count > len(nums)/2
}

// 通过双指针找到目标数字的左右边界，然后计算个数
func isMajorityElementIII(nums []int, target int) bool {
	if len(nums) == 1 {
		return nums[0] == target
	}

	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] < target {
			left++
		} else if nums[left] > target {
			return false
		}

		if nums[right] > target {
			right--
		} else if nums[right] < target {
			return false
		}

		if nums[left] == target && nums[right] == target {
			break
		}
	}

	return right-left+1 > len(nums)/2
}

// 二分查找找到目标数字的左右边界，然后计算个数
func isMajorityElementIV(nums []int, target int) bool {
	left, right := binarySearchLeft(nums, target), binarySearchRight(nums, target)
	return left != -1 && right != -1 && right-left+1 > len(nums)/2
}

func binarySearchLeft(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l < len(nums) && nums[l] == target {
		return l
	}
	return -1
}

func binarySearchRight(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if r >= 0 && nums[r] == target {
		return r
	}
	return -1
}

// 只需要一次二分查找
// 找到左边界，然后判断左边界加上n/2是否等于目标数字
func isMajorityElementV(nums []int, target int) bool {
	left := binarySearchLeft(nums, target)
	return left != -1 && left+len(nums)/2 < len(nums) && nums[left+len(nums)/2] == target
}