package leetcode

// 单调栈 + 贪心
func removeDuplicateLetters(s string) string {
	stack := []rune{}
	// 是否已在栈里
	seen := make(map[rune]bool)
	// 字符最后出现的位置
	lastIndex := make(map[rune]int)

	for i, c := range s {
		lastIndex[c] = i
	}

	for i, c := range s {
		if seen[c] {
			continue
		}

		// stack[len(stack)-1] > c ,栈顶元素大于当前元素
		// lastIndex[stack[len(stack)-1]] > i 且栈顶元素之后还会出现 (字典序最小)
		// 则弹出栈顶元素
		for len(stack) > 0 && stack[len(stack)-1] > c && lastIndex[stack[len(stack)-1]] > i {
			delete(seen, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, c)
		seen[c] = true
	}

	return string(stack)
}