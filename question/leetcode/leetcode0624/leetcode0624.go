package leetcode

func maxDistance(arrays [][]int) int {
	res := 0
	n := len(arrays[0])
	minVal := arrays[0][0]
	maxVal := arrays[0][n-1]
	for i := 1; i < len(arrays); i++ {
		n = len(arrays[i])
		res = max(res, max(int(arrays[i][n-1]-minVal), int(maxVal-arrays[i][0])))
		minVal = min(minVal, arrays[i][0])
		maxVal = max(maxVal, arrays[i][n-1])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}