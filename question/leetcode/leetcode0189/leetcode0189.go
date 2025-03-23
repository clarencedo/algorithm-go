package leetcode

// 解法一 时间复杂度 O(n)，空间复杂度 O(1)
// 三次翻转
// 1. 翻转整个数组
// 2. 翻转前 k 个元素
// 3. 翻转后 n-k 个元素
// 例如：[1, 2, 3, 4, 5, 6, 7]，k = 3
// 1. 翻转整个数组：[7, 6, 5, 4, 3, 2, 1]
// 2. 翻转前 k 个元素：[5, 6, 7, 4, 3, 2, 1]
// 3. 翻转后 n-k 个元素：[5, 6, 7, 1, 2, 3, 4]
// 结果：[5, 6, 7, 1, 2, 3, 4]
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// 解法二 时间复杂度 O(n)，空间复杂度 O(n)
func rotateII(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

// 解法三 时间复杂度 O(n)，空间复杂度 O(1)
// 原地替换，不需要额外空间
func rotateIII(nums []int, k int) {
	//用一个临时变量保存被替换的值，就不需要重新申请一个空间
	n := len(nums)
	k %= n
	count := 0
	for start := 0; count < n; start++ {
		currentIndex := start
		//保存待替换的值
		prev := nums[start]
		for {
			nextIndex := (currentIndex + k) % n
			tmp := nums[nextIndex]
			nums[nextIndex] = prev
			prev = tmp

			currentIndex = nextIndex
			count++
			// start != currentIndex就一直循环
			if start == currentIndex {
				break
			}
		}
	}
}