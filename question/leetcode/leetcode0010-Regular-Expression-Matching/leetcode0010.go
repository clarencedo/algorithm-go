package leetcode

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	// i, j 分别表示 s 和 p 的索引
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	//f[i][j]表示s的前i个字符和p的前j个字符是否匹配
	//即s[0:i]和p[0:j]是否匹配
	//f[0][0]表示s[0:0]和p[0:0]是否匹配，即空字符串是否匹配
	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true

	//动态转移方程
	//p[j-1]不是*号
	//f[i][j] = f[i-1][j-1] && matches(i, j)
	//p[j-1]是*号
	//f[i][j] = f[i][j-2] || (matches(i, j-1) && f[i-1][j])

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}
