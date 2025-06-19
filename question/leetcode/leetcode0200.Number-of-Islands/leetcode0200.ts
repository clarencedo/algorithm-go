function numIslands(grid: string[][]): number {
	if (grid.length === 0) {
		return 0
	}
	const rows = grid.length;
	const cols = grid[0].length;
	let count = 0;

	for (let i = 0; i < rows; i++) {
		for (let j = 0; j < cols; j++) {

			if (grid[i][j] === '1') {
				dfs(grid, i, j);
				count += 1;
			}
		}
	}

	return count;
};

function dfs(grid: string[][], row: number, col: number) {
	if (row < 0 || col < 0 || row >= grid.length || col >= grid[0].length || grid[row][col] !== '1') {
		return
	}

	grid[row][col] = '0'

	dfs(grid, row - 1, col)
	dfs(grid, row + 1, col)
	dfs(grid, row, col - 1)
	dfs(grid, row, col + 1)
}