class Solution:
    def climbStairs(self, n: int) -> int:
        if n <= 2:
            return n
        pre_pre, pre = 1, 2
        cur = 0
        for i in range(3, n+1):
            cur = pre_pre + pre
            pre_pre, pre = pre, cur

        return cur

