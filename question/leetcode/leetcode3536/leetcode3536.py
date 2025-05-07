class Solution:

    # 排序时间复杂度 O(kLogk)
    def maxProduct(self, n: int) -> int:
        s = sorted(str(n))
        return int(s[-1]) * int(s[-2])

    # O(k)
    def maxProductII(self, n: int) -> int:
        max1, max2 = 0, 0

        while n > 0:
            d = n % 10
            if d > max1:
                max1, max2 = d, max1
            elif d > max2:
                max2 = d
            n //= 10

        return max1 * max2
