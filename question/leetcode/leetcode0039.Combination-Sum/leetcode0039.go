package leetcode

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var track []int

	backtrack(0, candidates, target, track, &res)

	return res
}

func backtrack(start int, candidates []int, target int, track []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, append([]int{}, track...))
		return
	}

	for i := start; i < len(candidates); i++ {
		if target < candidates[i] {
			continue
		}
		track = append(track, candidates[i])
		backtrack(i, candidates, target-candidates[i], track, res)
		track = track[:len(track)-1]
	}
}