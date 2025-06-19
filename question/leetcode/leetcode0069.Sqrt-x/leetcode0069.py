import math


class Solution:
    def mySqrt(self, x: int) -> int:
        if x == 0:
            return 0
        ans = int(math.exp(0.5 * math.log(x)))
        return ans + 1 if (ans + 1) ** 2 <= x else ans

    def mySqrt(self, x: int) -> int:
        '''
        二分法
        mid * mid <= x
        (mid + 1) * (mid + 1) > x
        '''
        left, right, ans = 0, x, -1
        while left <= right:
            mid = (left + right) // 2
            if mid * mid <= x:
                ans = mid
                left = mid + 1
            else:
                right = mid - 1
        return ans
