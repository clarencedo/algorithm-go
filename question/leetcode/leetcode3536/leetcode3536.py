class Solution:

    # 排序时间复杂度 O(kLogk)
    def maxProduct(self, n: int) -> int:
        s = sorted(str(n))
        return int(s[-1]) * int(s[-2])

    # O(k)
    def maxProductII(self, n: int) -> int:
        '''
        1. 用 max1 和 max2 分别记录当前最大的数字和次大的数字
        2. 遍历 n 的每一位数字，如果当前数字大于 max1，则将 max2 赋值为 max1，然后将 max1 赋值为当前数字；否则，如果当前数字大于 max2，则将 max2 赋值为当前数字
        '''
        max1, max2 = 0, 0

        while n > 0:
            d = n % 10
            if d > max1:
                max1, max2 = d, max1
            elif d > max2:
                max2 = d
            n //= 10

        return max1 * max2
