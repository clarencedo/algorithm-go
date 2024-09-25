package leetcode

func isOneEditDistance(s string, t string) bool {
	len1, len2 := len(s), len(t)
	if len1 > len2 {
		return isOneEditDistance(t, s)
	}
	if len2-len1 > 1 {
		return false
	}
	i, j, diff := 0, 0, 0
	for i < len1 {
		if s[i] == t[j] {
			i++
			j++
			continue
		}
		if len1 == len2 && diff == 0 {
			diff++
			i++
			j++
			continue
		} else if len1 != len2 && diff == 0 && s[i] == t[j+1] {
			i++
			j += 2
			diff++
			continue
		} else {
			return false
		}
	}
	return diff == 1 || (len1 != len2 && diff == 0 && i == len1)
}