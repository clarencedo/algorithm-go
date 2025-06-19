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
            res = max(res, pre_sum[i] - pre_sum[dq[0]]
                      if dq else float('-inf'))

            # 维护单调递增队列（去掉比当前 pre_sum[i] 大的）
            while dq and pre_sum[dq[-1]] >= pre_sum[i]:
                dq.pop()

            dq.append(i)

        return res

    # 不跨越数组边界：即普通的最大子数组和问题（Kadane算法可解决）
    # 跨越数组边界：这种情况可以转化为"总和减去最小子数组和"

    def maxSubarraySumCircular(self, nums: List[int]) -> int:
        total = sum(nums)
        # cur_max 遍历到当前元素时，以当前元素结尾的子数组的最大和
        # max_sum 到目前为止，在非环形情况下找到的全局最大子数组和
        # cur_min 遍历到当前元素时，以当前元素结尾的子数组的最小和
        # min_sum 到目前为止，在非环形情况下找到的全局最小子数组和
        cur_max = max_sum = nums[0]
        cur_min = min_sum = nums[0]

        for num in nums[1:]:
            cur_max = max(num, cur_max + num)
            max_sum = max(max_sum, cur_max)

            cur_min = min(num, cur_min + num)
            min_sum = min(min_sum, cur_min)

        # 如果max_sum < 0，说明数组中没有正数，
        if max_sum < 0:
            return max_sum
        return max(max_sum, total - min_sum)
