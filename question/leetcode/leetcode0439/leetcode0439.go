package leetcode

func parseTernary(expression string) string {
	stack := []byte{}
	n := len(expression)

	for i := n - 1; i >= 0; i-- {
		ch := expression[i]

		if ch == '?' {
			if len(stack) < 2 || i-1 < 0 {
				return "Error: invalid expression"
			}

			trueCase := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			falseCase := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if expression[i-1] == 'T' {
				stack = append(stack, trueCase)
			} else if expression[i-1] == 'F' {
				stack = append(stack, falseCase)
			}
			i--
		} else if ch != ':' {
			stack = append(stack, ch)
		}
	}

	if len(stack) != 1 {
		return "Error: invalid expression"
	}
	return string(stack)
}