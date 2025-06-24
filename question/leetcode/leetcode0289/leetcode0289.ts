/**
 Do not return anything, modify board in-place instead.
 */
// TODO: 理解状态
function gameOfLife(board: number[][]): void {
	let m = board.length, n = board[0].length;

	const directions = [
		[-1, -1], [-1, 0], [-1, 1],
		[0, -1], [0, 1],
		[1, -1], [1, 0], [1, 1],
	];

	for (let i = 0; i < m; i++) {
		for (let j = 0; j < n; j++) {
			let liveNeighbors = 0;

			for (const [dx, dy] of directions) {
				const x = i + dx;
				const y = j + dy;

				if (x >= 0 && x < m && y >= 0 && y < n) {
					// 只看原状态
					liveNeighbors += board[x][y] % 2;
				}
			}

			if (board[i][j] === 1) {
				//活
				if (liveNeighbors === 2 || liveNeighbors === 3) {
					board[i][j] = 3; //活->活
				} else {
					board[i][j] = 1;//活->死
				}
			} else {
				if (liveNeighbors === 3) {
					board[i][j] = 2; // 死 ->活
				}
			}
		}
	}


	for (let i = 0; i < m; i++) {
		for (let j = 0; j < n; j++) {
			board[i][j] = board[i][j] >> 1; //取下一状态
		}
	}

};