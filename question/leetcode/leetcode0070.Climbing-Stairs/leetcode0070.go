package leetcode

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prePre, pre := 1, 2
	var curr int
	for i := 3; i <= n; i++ {
		curr = prePre + pre
		prePre, pre = pre, curr
	}

	return curr
}