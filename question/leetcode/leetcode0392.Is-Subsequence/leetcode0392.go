package leetcode

func isSubsequence(s string, t string) bool {
	p1, p2 := 0, 0
	n1, n2 := len(s), len(t)
	if n1 == 0 {
		return true
	}
	for p2 < n2 {
		if s[p1] == t[p2] {
			p1++
			if p1 == n1 {
				return true
			}
		}
		p2++
	}
	return false
}