import math


class Solution:
    def fib(self, n: int) -> int:
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

    def fib(self, n: int) -> int:
        if n < 2:
            return n
        dp = [0] * (n+1)
        dp[1] = 1
        for i in range(2, n+1):
            dp[i] = dp[i-1] + dp[i-2]

        return dp[n]

    def fib(self, n: int) -> int:
        sqrt5 = math.sqrt(5)
        phi = (1 + sqrt5) / 2
        return round(pow(phi, n) / sqrt5)
