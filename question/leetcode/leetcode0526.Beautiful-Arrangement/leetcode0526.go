package leetcode

func countArrangement(n int) int {
	// 记录数字x 是否已使用
	visit := make([]bool, n+1)
	// 每个位置i的合法的数字列表
	match := make([][]int, n+1)
	ans := 0
	//初始化每个位置能放哪些数字
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i%j == 0 || j%i == 0 {
				match[i] = append(match[i], j)
			}
		}
	}
	var backtrack func(int)
	backtrack = func(index int) {
		if index > n {
			//当前填完了所有的位置，找到一个合法排列
			ans++
			return
		}

		//枚举当前i位置能放的合法数字
		for _, x := range match[index] {
			//如果没有使用过,则尝试使用
			if !visit[x] {
				visit[x] = true
				backtrack(index + 1)
				visit[x] = false
			}
		}
	}

	backtrack(1)
	return ans
}