package leetcode

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
func longestCommonPrefixII(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	prefix := ""
	firstStr := strs[0]
	for i := 0; i < len(firstStr); i++ {
		char := firstStr[i]
		for j := 1; j < len(strs); j++ {
			//防止索引越界
			if i >= len(strs[j]) || strs[j][i] != char {
				return prefix
			}
		}
		prefix += string(char)
	}

	return prefix
}