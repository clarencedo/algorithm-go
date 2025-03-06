package leetcode

// TODO: 重新理解
func findPermutation(s string) []int {
	n := len(s)
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = i + 1
	}
	for i := 0; i < n; i++ {
		if s[i] == 'D' {
			j := i
			for j < n && s[j] == 'D' {
				j++
			}
			reverse(ans, i, j)
			i = j
		}
	}
	return ans
}

func reverse(ans []int, i, j int) {
	for i < j {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
}