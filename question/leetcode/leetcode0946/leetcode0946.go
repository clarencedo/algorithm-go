package leetcode

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	i := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}
	return len(stack) == 0
}
