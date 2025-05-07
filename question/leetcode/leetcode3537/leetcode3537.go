package leetcode

func specialGrid(n int) [][]int {
	val := 0
	var dfs func([][]int, int, int)
	dfs = func(a [][]int, l, r int) {
		if len(a) == 1 {
			a[0][l] = val
			val++
			return
		}
		m := len(a) / 2
		dfs(a[:m], l+m, r)
		dfs(a[m:], l+m, r)
		dfs(a[m:], l, l+m)
		dfs(a[:m], l, l+m)
	}

	a := make([][]int, 1<<n)
	for i := range a {
		a[i] = make([]int, 1<<n)
	}

	dfs(a, 0, 1<<n)
	return a
}
