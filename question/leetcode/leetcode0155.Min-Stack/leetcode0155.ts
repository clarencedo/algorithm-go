class MinStack {
	private stack: number[]; //正常栈
	private minStack: number[]; //最小栈

	constructor() {
		this.stack = [];
		this.minStack = [];
	}

	push(val: number): void {
		this.stack.push(val)
		// 如果minStafck为空，或当前值更小，则入辅助栈
		if (this.minStack.length === 0 || val <= this.minStack[this.minStack.length - 1]) {
			this.minStack.push(val);
		}
	}

	pop(): void {
		if (this.stack.length > 0) {
			const val = this.stack.pop();
			// 如果弹出的是最小值，则也要minStack出栈
			if (val === this.minStack[this.minStack.length - 1]) {
				this.minStack.pop();
			}

		}
	}

	top(): number {
		return this.stack[this.stack.length - 1];

	}

	getMin(): number {
		return this.minStack[this.minStack.length - 1];
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * var obj = new MinStack()
 * obj.push(val)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */


