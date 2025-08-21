package leetcode

// 双指针： 对于每个i，把j尽量右推到满足 nums1[i]<=nums2[j],更新答案
func maxDistance(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	ans := 0
	j := 0
	for i := range n1 {
		j = max(j, i)
		for j+1 < n2 && nums2[j+1] >= nums1[i] {
			j++
		}
		if nums2[j] > nums1[i] {
			ans = max(ans, j-i)
		}

	}
	return ans
}

// 二分枚举 nums1[i] 在 nums2 中能到达的最右 j，并取最大 j-i
func maxDistanceII(nums1, nums2 []int) int {
	ans := 0
	for i, val := range nums1 {
		j := rightmostGEInNonIncreasing(nums2, val)
		if j != -1 && j >= i {
			if j-i > ans {
				ans = j - i
			}
		}
	}
	return ans
}

func rightmostGEInNonIncreasing(arr []int, target int) int {
	lo, hi := 0, len(arr)-1
	pos := -1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if arr[mid] >= target { // 条件成立，记录并尽量往右找更大的 j
			pos = mid
			lo = mid + 1
		} else { // arr[mid] < target，右边更小，应该往左缩小 hi
			hi = mid - 1
		}
	}
	return pos
}
