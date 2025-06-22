package leetcode

import "strings"

// 方法一： 单调栈
// 时间复杂度 O(n)
func removeKdigits(nums string, k int) string {
	n := len(nums)
	stack := []byte{}
	for i := 0; i < n; i++ {

		// 栈顶元素大于当前元素，则移出栈顶元素
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > nums[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, nums[i])
	}
	// 结束遍历时 可能还有删除次数，或者是最左以0开头
	// 如果 k > 0，说明还有删除次数，继续删除
	stack = stack[:len(stack)-k]
	// 删除掉前导0
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}

	return ans
}
