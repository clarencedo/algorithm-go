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

func validateStackSequencesII(pushed []int, popped []int) bool {
	//初始化长度为0，长度作为判断条件
	stack := make([]int, 0)

	//在popped中的元素下标
	j := 0
	for i := 0; i < len(pushed)-1; i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}

	return len(stack) == 0
}
