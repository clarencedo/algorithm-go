package leetcode

// TODO: 重新理解旋转
func rotate(matrix [][]int) {
	// 获取矩阵的行数，假设矩阵为n x n
	n := len(matrix)
	// 遍历矩阵的一半，因为在旋转过程中，每次操作会影响到四个元素
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			// 临时存储当前元素
			tmp := matrix[i][j]
			// 进行元素旋转操作
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = tmp
		}
	}
}