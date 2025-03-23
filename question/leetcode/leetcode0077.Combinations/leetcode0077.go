package leetcode

func combine(n int, k int) [][]int {
	var res [][]int
	var track []int
	var backtrack func(int)
	backtrack = func(start int) {
		if len(track) == k {
			tmp := make([]int, k)
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for i := start; i <= n; i++ {
			track = append(track, i)
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtrack(1)
	return res
}

func combine2(n int, k int) [][]int {
	var res [][]int
	var track []int
	backtrack2(1, n, k, track, &res)
	return res
}

func backtrack2(start, n, k int, track []int, res *[][]int) {
	if len(track) == k {
		tmp := make([]int, k)
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}

	for i := start; i <= n; i++ {
		track = append(track, i)
		backtrack2(i+1, n, k, track, res)
		track = track[:len(track)-1]
	}
}

func combineII(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	generateCombinations(n, k, 1, c, &res)
	return res
}

func generateCombinations(n, k, start int, c []int, res *[][]int) {
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	// i will at most be n - (k - c.size()) + 1
	for i := start; i <= n-(k-len(c))+1; i++ {
		c = append(c, i)
		generateCombinations(n, k, i+1, c, res)
		c = c[:len(c)-1]
	}
	return
}
