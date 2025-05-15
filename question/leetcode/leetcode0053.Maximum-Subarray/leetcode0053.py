from collections import deque


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        ans = nums[0]

        for i in range(1, len(nums)):
            if nums[i-1] + nums[i] > nums[i]:
                nums[i] = nums[i-1] + nums[i]

            if nums[i] > ans:
                ans = nums[i]

        return ans

    # 数组带环，Leetcode918
    def maxSubArrayCircular(self, nums: List[int]) -> int:
        n = len(nums)
        # 构建前缀和数组， pre[i]表示nums[0]到nums[i-1]的和
        pre_sum = [0] * (2 * n + 1)
        for i in range(1, 2 * n+1):
            pre_sum[i] = pre_sum[i-1] + nums[(i-1) % n]

        # 设置为负无穷
        res = float('-inf')
        # dq保存的是pre_sum的下标，且保持单调递增
        dq = deque()

        for i in range(1, 2 * n + 1):
            if dq and dq[0] < i - n:
                dq.popleft()

            res = max(res, pre_sum[i] - pre_sum[dp[0]]
                      if dp else float('-inf'))

            while dq and pre_sum[i] >= pre_sum[dq[-1]]:
                dq.pop()

            dq.append(i)

        return res
