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

func lengthOfLongestSubstringIII(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	lastIndex := make(map[rune]int)
	maxLen := 0

	for i, v := range s {
		if i == 0 {
			dp[0] = 1
		} else {
			// 当前字符s[i] 是否出现在以i-1结尾的最长无重复子串中
			// dp[i-1] 是上一个位置能形成的不重复子串的长度
			// 所以这个子串的起始位置是 i-dp[i]
			// 也就是说，以s[i-1]结尾的不重复子串是: s[i-dp[i-1]] , ..., s[i-1]
			// 如果last上一次这个字符出现的位置 早于，子串的起始为止
			// 那么当前字符s[i]，在s[i-dp[i-1]],...,s[i-1]中不重复
			if last, ok := lastIndex[v]; !ok || last < i-dp[i-1] {
				dp[i] = dp[i-1] + 1
			} else {
				dp[i] = i - last
			}
		}
		lastIndex[v] = i
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	return maxLen
}