function largestIsland(grid: number[][]): number {
	const n = grid.length;
	const areaMap = new Map<number, number>(); // islandId -> area
	let islandId = 2; // 从 2 开始标记岛屿，避免和 0,1 冲突

	// 四个方向
	const dirs = [[0, 1], [1, 0], [0, -1], [-1, 0]];

	// 深度优先搜索，标记岛屿编号，并返回面积
	function dfs(x: number, y: number, id: number): number {
		if (x < 0 || x >= n || y < 0 || y >= n || grid[x][y] !== 1) return 0;

		grid[x][y] = id;
		let area = 1;

		for (const [dx, dy] of dirs) {
			area += dfs(x + dx, y + dy, id);
		}
		return area;
	}

	// 第一步：给所有岛屿标号，并记录面积
	for (let i = 0; i < n; ++i) {
		for (let j = 0; j < n; ++j) {
			if (grid[i][j] === 1) {
				const area = dfs(i, j, islandId);
				areaMap.set(islandId, area);
				islandId++;
			}
		}
	}

	let maxArea = Math.max(...areaMap.values(), 0); // 如果全是 1

	// 第二步：尝试把每一个 0 翻成 1，计算连通面积
	for (let i = 0; i < n; ++i) {
		for (let j = 0; j < n; ++j) {
			if (grid[i][j] === 0) {
				const seen = new Set<number>();

				for (const [dx, dy] of dirs) {
					const ni = i + dx;
					const nj = j + dy;
					if (ni >= 0 && ni < n && nj >= 0 && nj < n) {
						const id = grid[ni][nj];
						if (id > 1) {
							seen.add(id);
						}
					}
				}

				let area = 1; // 翻转自身的 1
				for (const id of seen) {
					area += areaMap.get(id)!;
				}
				maxArea = Math.max(maxArea, area);
			}
		}
	}

	return maxArea;
}

//超时的解法
function reverseIsland(grid: number[][]): number {
	const rows = grid.length;
	const cols = grid[0].length;
	let maxArea = 0;

	// 深拷贝 grid
	const copyGrid = (grid: number[][]) =>
		grid.map(row => [...row]);

	// 计算岛屿面积（连通块），包含已翻转为 2 的位置
	const findMax = (grid: number[][], i: number, j: number, visited: boolean[][]): number => {
		if (
			i < 0 || i >= rows || j < 0 || j >= cols ||
			visited[i][j] || (grid[i][j] !== 1 && grid[i][j] !== 2)
		) {
			return 0;
		}

		visited[i][j] = true;
		let area = 1;

		const directions = [[0, 1], [1, 0], [-1, 0], [0, -1]];
		for (const [dx, dy] of directions) {
			area += findMax(grid, i + dx, j + dy, visited);
		}

		return area;
	};

	// 尝试翻转某个 0，并计算可能最大连通面积
	const tryReverse = (i: number, j: number): number => {
		const newGrid = copyGrid(grid);
		newGrid[i][j] = 2; // 临时翻转
		const visited: boolean[][] = Array.from({ length: rows }, () => Array(cols).fill(false));
		return findMax(newGrid, i, j, visited);
	};

	let hasZero = false;

	// 遍历所有的 0，尝试翻转
	for (let i = 0; i < rows; i++) {
		for (let j = 0; j < cols; j++) {
			if (grid[i][j] === 0) {
				hasZero = true;
				const area = tryReverse(i, j);
				maxArea = Math.max(maxArea, area);
			}
		}
	}

	// 如果没有任何 0，说明已经是最大岛，返回全图面积
	if (!hasZero) {
		const visited: boolean[][] = Array.from({ length: rows }, () => Array(cols).fill(false));
		return findMax(grid, 0, 0, visited);
	}

	return maxArea;
}