function dailyTemperatures(temperatures: number[]): number[] {
	const res: number[] = new Array(temperatures.length).fill(0);
	// 单调栈存储索引,单调递减，栈顶元素温度最小
	const stack: number[] = [];

	for (let i = 0; i < temperatures.length; i++) {
		//如果当前元素 大于栈顶元素索引位置的值，则出栈计算 栈顶元素索引位置 res里该放的值
		while (stack.length > 0 && temperatures[i] > temperatures[stack[stack.length - 1]]) {
			const stackTop = stack.pop()!;
			res[stackTop] = i - stackTop;
		}
		stack.push(i);
	}

	return res;
};