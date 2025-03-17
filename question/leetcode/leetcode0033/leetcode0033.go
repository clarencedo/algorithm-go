package leetcode

func search(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	if n == 0 {
		return -1
	}

	if n == 1 && nums[0] == target {
		return 0
	}

	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			//左边有序
			if nums[0] <= target && target < nums[mid] {
				// target在左边
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			//右边有序
			if nums[mid] < target && target <= nums[n-1] {
				// target在右边
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}