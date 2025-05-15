from collections import deque
from typing import List

class Solution:
    def maxSubarraySumCircular(self, nums: List[int]) -> int:
        n = len(nums)
        # 构建前缀和数组：pre[i] 表示 nums[0] 到 nums[i-1] 的前缀和
        pre_sum = [0] * (2 * n + 1)
        for i in range(1, 2 * n + 1):
            pre_sum[i] = pre_sum[i - 1] + nums[(i - 1) % n]

        res = float('-inf')
        dq = deque()  # 保存的是 pre_sum 的下标，且保持单调递增

        for i in range(1, 2 * n + 1):
            # 保证子数组长度不超过 n（滑动窗口大小最多为 n）
            if dq and dq[0] < i - n:
                dq.popleft()

            # 当前前缀和减去最小的前缀和，即是当前位置的最大子数组和
            res = max(res, pre_sum[i] - pre_sum[dq[0]] if dq else float('-inf'))

            # 维护单调递增队列（去掉比当前 pre_sum[i] 大的）
            while dq and pre_sum[dq[-1]] >= pre_sum[i]:
                dq.pop()

            dq.append(i)

        return res