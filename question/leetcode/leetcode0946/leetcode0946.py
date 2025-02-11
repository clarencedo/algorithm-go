from typing import List


class Solution:
    def validateStackSequences(pushed: List[int], popped: List[int]) -> bool:
        stack = []
        i = 0
        for v in pushed:
            stack.append(v)
            while stack and i < len(popped) and stack[-1] == popped[i]:
                stack.pop()
                i += 1
        return not stack