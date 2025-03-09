package leetcode

// 递归 + 记忆化搜索
// 如果当前状态下，有两个连续的'++'，那么可以将其变为'--'
// 然后递归调用dfs函数，如果返回false，说明当前状态可以赢
func canWin(currentState string) bool {
	memo := make(map[string]bool)
	return dfs(currentState, memo)
}

func dfs(currentState string, memo map[string]bool) bool {
	if v, ok := memo[currentState]; ok {
		return v
	}
	for i := 0; i < len(currentState)-1; i++ {
		if currentState[i] == '+' && currentState[i+1] == '+' {
			nextState := currentState[:i] + "--" + currentState[i+2:]
			if !dfs(nextState, memo) {
				memo[currentState] = true
				return true
			}
		}
	}
	memo[currentState] = false
	return false
}

// 动态规划
// dp[i]表示当前状态下，是否是必胜态
// dp[i] =any()