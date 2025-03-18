package leetcode

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	max := 0
	//hash存储字符出现的次数
	hash := make([]int, 256)
	//i为窗口左边界，j为窗口右边界
	for i, j := 0, -1; j < n-1; {
		j++
		hash[s[j]]++
		for hash[s[j]] > 1 {
			hash[s[i]]--
			i++
		}
		if j-i+1 > max {
			max = j - i + 1
		}
	}
	return max
}

// 滑动窗口另一种写法
func lengthOfLongestSubstringII(s string) int {
	window := make(map[byte]int)
	left, right := 0, 0
	res := 0
	for right < len(s) {
		rightChar := s[right]
		window[rightChar]++
		right++

		for window[rightChar] > 1 {
			leftChar := s[left]
			window[leftChar]--
			left++
		}
		res = max(res, right-left)
	}
	return res
}