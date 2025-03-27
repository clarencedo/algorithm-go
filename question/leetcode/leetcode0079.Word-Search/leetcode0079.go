package leetcode

// 时间复杂度：O(3^k * mn)，其中 k 为字符串 word 的长度，即需要遍历 3^k 个节点。
// 空间复杂度：O(k)，其中 k 为字符串 word 的长度，即递归调用的栈深度。
func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	var backtrack func(int, int, int) bool
	backtrack = func(i, j, k int) bool {
		// k 为当前匹配的字符索引
		if i >= n || i < 0 || j >= m || j < 0 || board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		temp := board[i][j]
		// 标记已访问
		// 由于题目要求不能重复使用，所以这里直接修改为特殊字符, 0随便用的，只要不是字母就行
		board[i][j] = 0
		res := backtrack(i+1, j, k+1) || backtrack(i-1, j, k+1) || backtrack(i, j+1, k+1) || backtrack(i, j-1, k+1)
		// 恢复原值
		board[i][j] = temp
		return res
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}
	return false
}

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func existII(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i, v := range board {
		for j := range v {
			if searchWord(board, visited, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func isInBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func searchWord(board [][]byte, visited [][]bool, word string, index, x, y int) bool {
	if index == len(word)-1 {
		return board[x][y] == word[index]
	}
	if board[x][y] == word[index] {
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			nx := x + dir[i][0]
			ny := y + dir[i][1]
			if isInBoard(board, nx, ny) && !visited[nx][ny] && searchWord(board, visited, word, index+1, nx, ny) {
				return true
			}
		}
		visited[x][y] = false
	}
	return false
}
