package leetcode

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	res := []int{}
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	for len(res) < len(matrix)*len(matrix[0]) {
		// 从左到右遍历上边界
		for i := left; i <= right && top <= bottom; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		// 从上到下遍历右边界
		for i := top; i <= bottom && left <= right; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		// 从右到左遍历下边���
		for i := right; i >= left && top <= bottom; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--

		// 从下到上遍历左边界
		for i := bottom; i >= top && left <= right; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}