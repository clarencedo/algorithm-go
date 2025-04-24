package leetcode

// 用第一行和第一列当标记位，因为他们只有有一个0，那么他们其他都会变成0
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row0HasZero, col0HasZero := false, false

	//检查第一列是否有0
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0HasZero = true
			break
		}
	}

	//检查第一行是否有0
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			row0HasZero = true
			break
		}
	}

	//用第一行和第一列座位标记位
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	//根据标记清零，不处理第一行和第一列
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	//最后处理第一行和第一列
	if row0HasZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if col0HasZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
