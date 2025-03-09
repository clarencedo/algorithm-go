package leetcode

// values[j] + values[j] + i - j = values[i] + i + values[j] - j
// 拆分为两个部分，values[j] + i 和 values[j] - j
// mx = values[j] + i
func maxScoreSjghtseeingPair(values []int) int {
	n := len(values)
	ans, mx := 0, values[0]
	for j := 1; j < n; j++ {
		ans = max(ans, mx+values[j]-j)
		mx = max(mx, values[j]+j)
	}
	return ans
}