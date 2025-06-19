import math


class Solution:
    # 递归方式实现（简单但效率低）
    def fib_recursive(self, n: int) -> int:
        if n <= 0:
            return 0
        if n == 1:
            return 1
        return self.fib_recursive(n-1) + self.fib_recursive(n-2)

    def fib(self, n: int) -> int:
        if n <= 0:
            return 0
        if n == 1:
            return 1
        return self.fib(n-1) + self.fib(n-2)

    # 带备忘录的递归方式（自顶向下）

    def fib_memo(self, n: int) -> int:
        memo = [-1] * (n+1)

        def dp(n):
            if n <= 1:
                return n
            if memo[n] != -1:
                return memo[n]
            memo[n] = dp(n-1) + dp(n-2)
            return memo[n]

        return dp(n)

    # 迭代方式（空间优化）
    def fib_iter(self, n: int) -> int:
        if n < 2:
            return n

        # p -> f(n-2)
        # q -> f(n-1)
        # r -> f(n)
        p, q, r = 0, 0, 1
        for i in range(2, n + 1):
            p, q = q, r
            r = p + q

        return r

    # 动态规划方式（自底向上）
    def fib_dp(self, n: int) -> int:
        if n < 2:
            return n
        dp = [0] * (n+1)
        dp[1] = 1
        for i in range(2, n+1):
            dp[i] = dp[i-1] + dp[i-2]

        return dp[n]

    # 数学公式方式（黄金分割比）
    def fib_math(self, n: int) -> int:
        sqrt5 = math.sqrt(5)
        phi = (1 + sqrt5) / 2
        return round(pow(phi, n) / sqrt5)
