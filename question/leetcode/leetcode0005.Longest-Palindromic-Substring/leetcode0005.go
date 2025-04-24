package leetcode

// 解法一：中心扩散法, 时间复杂度O(n^2), 空间复杂度O(1)
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)   //奇数长度回文, 以i为中心
		len2 := expandAroundCenter(s, i, i+1) //偶数长度回文, 以i和i+1为中心
		len := max(len1, len2)
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// 解法二：动态规划 , 时间复杂度O(n^2), 空间复杂度O(n^2)
func longestPalindromeII(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	//单个字符都是回文
	for i := range dp {
		dp[i][i] = true
	}

	start, maxLength := 0, 1

	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			// j为子串的右边界
			j := i + l - 1
			//子串长度为2或3时，只需判断首尾字符是否相等,特殊处理
			if l == 2 || l == 3 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1] //一个回文串去掉首尾字符仍是回文串
			}

			if dp[i][j] && l > maxLength {
				start = i
				maxLength = l
			}
		}
	}
	return s[start : start+maxLength]

}

// 解法三：滑动窗口，时间复杂度O(n^2), 空间复杂度O(1)
func longestPalindromeIII(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	// 从最大可能的窗口开始
	for windowSize := n; windowSize > 0; windowSize-- {
		// 滑动窗口
		for start := 0; start+windowSize <= n; start++ {
			substr := s[start : start+windowSize]
			if isPalindrome(substr) {
				return substr
			}
		}
	}

	return string(s[0]) // 至少返回单个字符
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}