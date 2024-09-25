package leetcode

// slide window
func numKLenSubstrNoRepeats(s string, k int) int {
	n := len(s)
	left, right := 0, 0
	window := make(map[byte]int)
	res := 0

	for right < n {
		rightChar := s[right]
		window[rightChar]++
		right++
		for right-left > k || window[rightChar] > 1 {
			leftChar := s[left]
			window[leftChar]--
			left++
		}
		if right-left == k {
			res++
		}
	}

	return res
}