package main

func canMakeEqual(nums []int, k int) bool {
	return tryTarget(nums, 1, k) || tryTarget(nums, -1, k)
}

func tryTarget(nums []int, target int, k int) bool {
	arr := make([]int, len(nums))
	copy(arr, nums)
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != target {
			arr[i] *= -1
			arr[i+1] *= -1
			count++
			if count > k {
				return false
			}
		}
	}
	// 最后一位不能翻转
	return arr[len(arr)-1] == target
}
