package leetcode

import "strings"

func reverseWords(s string) string {
	strs := strings.Fields(s)
	for left, right := 0, len(strs)-1; left < right; left, right = left+1, right-1 {
		strs[left], strs[right] = strs[right], strs[left]
	}

	return strings.Join(strs, " ")
}
