package leetcode

// 从右上角开始搜索,时间复杂度O(m+n),空间复杂度O(1)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row, col := 0, len(matrix[0])-1

	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			row++
		} else {
			col--
		}
	}
	return false
}

func searchMatrixII(matrix [][]int, target int) bool {
	i, j := len(matrix)-1, 0
	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}

	return false
}

// 逐行二分查找 时间复杂度O(mlog(n)),空间复杂度O(1)
func searchMatrixIII(matrix [][]int, target int) bool {
	for _, row := range matrix {
		left, right := 0, len(row)-1
		for left <= right {
			mid := left + (right-left)/2
			if row[mid] == target {
				return true
			} else if row[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}