package leetcode

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}

	return count
}

func dfs(grid [][]byte, row, col int) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] != '1' {
		return
	}

	grid[row][col] = '0'

	dfs(grid, row-1, col)
	dfs(grid, row+1, col)
	dfs(grid, row, col-1)
	dfs(grid, row, col+1)
}

// BFS
func numIslandsII(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	count := 0
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++
				queue := [][]int{{i, j}}
				grid[i][j] = '0'

				for len(queue) > 0 {
					cur := queue[0]
					queue = queue[1:]

					for _, dir := range directions {
						x, y := cur[0]+dir[0], cur[1]+dir[1]
						if x >= 0 && x < rows && y >= 0 && y < cols && grid[x][y] == '1' {
							queue = append(queue, []int{x, y})
							grid[x][y] = '0'
						}
					}
				}
			}

		}
	}
	return count
}
