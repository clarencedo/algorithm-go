package leetcode

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, 0)

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			stackTop := stack[len(stack)-1]
			res[stackTop] = i - stackTop
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}