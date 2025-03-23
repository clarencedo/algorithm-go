package leetcode

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !isalnum(s[left]) {
			left++
			continue
		}
		if !isalnum(s[right]) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 判断是否是字母或数字
func isalnum(ch byte) bool {
	return (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}