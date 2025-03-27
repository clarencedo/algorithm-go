package leetcode

func generateParenthesis(n int) []string {
	var res []string
	backtrack("", n, 0, 0, &res)
	return res
}

func backtrack(current string, n int, open, close int, res *[]string) {
	if open == n && close == n {
		*res = append(*res, current)
		return
	}

	if open < n {
		backtrack(current+"(", n, open+1, close, res)
	}
	if close < open {
		backtrack(current+")", n, open, close+1, res)
	}

}