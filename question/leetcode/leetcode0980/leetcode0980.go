package leetcode

const (
	ENDING = 2
	OBS    = -1
	START  = 1
)

func uniquePathsIII(grid [][]int) int {
	// 1. 找到起点和需要访问的空格总数
	startR, startC, empty := 0, 0, 1 // 包含起点所以初始为1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == START {
				startR, startC = i, j
			} else if grid[i][j] == 0 {
				empty++
			}
		}
	}

	// 2. 定义四个方向
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	count := 0

	// 3. 回溯函数
	var backtrack func(r, c, remain int)
	backtrack = func(r, c, remain int) {
		// 越��检查
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] < 0 {
			return
		}

		// 到达终点且所有空格已访问
		if grid[r][c] == ENDING {
			if remain == 0 {
				count++
			}
			return
		}

		// 标记���已访问
		grid[r][c] = -2
		remain--

		// 四个方向探索
		for _, dir := range dirs {
			backtrack(r+dir[0], c+dir[1], remain)
		}

		// 回溯
		grid[r][c] = 0
		remain++
	}

	// 4. 从起点开始回溯
	backtrack(startR, startC, empty)
	return count
}