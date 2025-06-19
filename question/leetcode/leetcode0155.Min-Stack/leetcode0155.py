class MinStack:
    def __init__(self):
        self.stack = []  # 正常栈
        self.min_stack = []  # 辅助栈，保存当前最小值

    def push(self, val: int) -> None:
        self.stack.append(val)
        # 如果 min_stack 为空 或 当前值更小，则入辅助栈
        if not self.min_stack or val <= self.min_stack[-1]:
            self.min_stack.append(val)

    def pop(self) -> None:
        if self.stack:
            val = self.stack.pop()
            # 如果弹出的是最小值，也要弹出 min_stack 中的
            if val == self.min_stack[-1]:
                self.min_stack.pop()

    def top(self) -> int:
        return self.stack[-1]

    def getMin(self) -> int:
        return self.min_stack[-1]