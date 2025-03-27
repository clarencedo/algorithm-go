/**
 * @param {number[][]} grid
 * @return {number}
 */
const uniquePathsIII = function (grid) {
	const ROWS = grid.length;
	const COLS = grid[0].length;
	const END = 2;
	const START = 1;
	const OBSTACLE = -1;
	const EMPTY = 0;

	let startR, startC, empty = 1; // 包含起点所以初始为1
	let count = 0;

	// 1. 找到起点和统计需要走过的空格数量
	for (let r = 0; r < ROWS; r++) {
		for (let c = 0; c < COLS; c++) {
			if (grid[r][c] === START) {
				startR = r;
				startC = c;
			} else if (grid[r][c] === EMPTY) {
				empty++;
			}
		}
	}

	// 2. 定义四个移动方向
	const dirs = [[0, 1], [1, 0], [0, -1], [-1, 0]];

	// 3. 回溯函数
	const backtrack = (r, c, remain) => {
		// 边界检查
		if (r < 0 || r >= ROWS || c < 0 || c >= COLS || grid[r][c] < 0) {
			return;
		}

		// 到达终点且所有空格已访问
		if (grid[r][c] === END) {
			if (remain === 0) count++;
			return;
		}

		// 标记为已访问（临时设为-2）
		grid[r][c] = -2;

		// 向四个方向探索
		for (const [dr, dc] of dirs) {
			backtrack(r + dr, c + dc, remain - 1);
		}

		// 回溯，恢复状态
		grid[r][c] = EMPTY;
	};

	// 4. 从起点开始回溯
	backtrack(startR, startC, empty);

	return count;
};