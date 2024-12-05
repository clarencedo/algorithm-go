package leetcode

// 倒序遍历，非':'字符 入栈，遇到'?'字符，弹出两个元素，根据'?'前的字符，选择入栈
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

// 递归
func parseTernaryII(expression string) string {
	// 递归解析三元表达式
	var helper func(expr string, start, end int) string
	helper = func(expr string, start, end int) string {
		// 如果表达式中没有 '?'，直接返回结果
		if start == end {
			return string(expr[start])
		}

		// 找到当前表达式的第一个 '?' 和对应的 ':'
		questionMark := start
		for expr[questionMark] != '?' {
			questionMark++
		}

		// 确定条件
		condition := expr[questionMark-1]

		// 找到与 '?' 对应的 ':'
		colon := questionMark + 1
		balance := 1 // 用于匹配嵌套的三元运算符
		for balance > 0 {
			if expr[colon] == '?' {
				balance++
			} else if expr[colon] == ':' {
				balance--
			}
			colon++
		}
		colon-- // 退回到真正的 ':'

		// 根据条件递归解析
		if condition == 'T' {
			return helper(expr, questionMark+1, colon-1) // 解析 trueCase
		}
		return helper(expr, colon+1, end) // 解析 falseCase
	}

	// 从完整表达式开始递归解析
	return helper(expression, 0, len(expression)-1)
}

func parseTernaryIII(expression string) string {
	return helper(expression, 0, len(expression)-1)
}

func helper(expression string, start, end int) string {
	// base case
	if start == end {
		return string(expression[start])
	}

	questionMark := start
	for expression[questionMark] != '?' {
		questionMark++
	}

	condition := expression[questionMark-1]

	colon := questionMark + 1
	balance := 1
	for balance > 0 {
		if expression[colon] == '?' {
			balance++
		} else if expression[colon] == ':' {
			balance--
		}
		colon++
	}
	colon--

	if condition == 'T' {
		return helper(expression, questionMark+1, colon-1)
	}
	return helper(expression, colon+1, end)
}