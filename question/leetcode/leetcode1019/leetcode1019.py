from typing import List, Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def nextLargerNodes(head: Optional[ListNode]) -> List[int]:
    res = []
    stack = []  # 存储 (值, 索引) 元组
    cur = head
    idx = -1

    while cur:
        idx += 1
        res.append(0)  # 预占位
        # 维护单调栈，栈顶元素小于当前值时出栈
        while stack and stack[-1][0] < cur.val:
            val, i = stack.pop()
            res[i] = cur.val  # 赋值下一个更大的值
        stack.append((cur.val, idx))
        cur = cur.next

    return res