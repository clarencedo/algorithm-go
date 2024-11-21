package leetcode

func missingElement(nums []int, k int) int {
	n := len(nums)

	//定义一个辅助函数， 计算前i个元素缺失的数字数量
	missing := func(i int) int {
		return nums[i] - nums[0] - i
	}

	//如果第k个缺失数字超过数组范围，直接计算结果
	if k > missing(n-1) {
		return nums[n-1] + (k - missing(n-1))
	}

	//二分查找，找到最后满足missing(mid) < k 的位置
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if missing(mid) < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left-1] + (k - missing(left-1))

}
