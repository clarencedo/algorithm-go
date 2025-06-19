package leetcode

// 用一个滑动窗口，窗口的大小是p的长度，然后窗口每次向右移动一位，，检查窗口内的字符是否与p构成变位词
// 用两个哈希表分别表示p和窗口内字符的频次，然后比较两个哈希表是否相等
func findAnagrams(s string, p string) []int {
	window := make(map[byte]int)
	pCount := make(map[byte]int)
	res := []int{}
	for _, char := range p {
		pCount[byte(char)]++
	}
	left, right := 0, 0
	//滑动窗口
	for right < len(s) {
		rightChar := s[right]
		window[rightChar]++
		right++

		//窗口大小等于p的长度
		if right-left == len(p) {
			//窗口内字符频次与p相等
			if isMatch(window, pCount) {
				res = append(res, left)
			}
			leftChar := s[left]

			if window[leftChar] > 1 {
				window[leftChar]--
			} else {
				delete(window, leftChar)
			}
			left++
		}

	}
	return res
}

func isMatch(a, b map[byte]int) bool {
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	for k, v := range b {
		if a[k] != v {
			return false
		}
	}
	return true
}
