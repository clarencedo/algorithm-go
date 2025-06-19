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

    # 前缀和
    def maxSubArII(self, nums: List[int]) -> int:
        if not nums:
            return 0

        # 前缀和数组，pre_sum[i]表示nums[0]到nums[i-1]的和
        pre_sum = [0] * (len(nums) + 1)
        for i in range(1, len(nums) + 1):
            pre_sum[i] = pre_sum[i-1] + nums[i-1]

        ans = nums[0]  # 初始化为第一个元素
        # 存储的是 pre_sum 的下标，且保持单调递增
        dq = deque([0])  # 初始化队列，放入前缀和数组的第一个元素的索引

        for i in range(1, len(pre_sum)):
            # 计算当前子数组的和
            ans = max(ans, pre_sum[i] - pre_sum[dq[0]])

            # 维护单调队列，保证队列中的元素对应的前缀和是单调递增的
            # 如果新的前缀和小于队列尾部的前缀和，则队列尾部的元素不可能是最优解
            while dq and pre_sum[dq[-1]] >= pre_sum[i]:
                dq.pop()

            dq.append(i)

        return ans

    # 数组带环，Leetcode918

    def maxSubArrayCircular(self, nums: List[int]) -> int:
        n = len(nums)
        # 构建前缀和数组， pre[i]表示nums[0]到nums[i-1]的和
        pre_sum = [0] * (2 * n + 1)
        for i in range(1, 2 * n+1):
            pre_sum[i] = pre_sum[i-1] + nums[(i-1) % n]

        res = float('-inf')
        # dq保存的是pre_sum的下标，且保持单调递增
        dq = deque([0])  # 初始化队列

        for i in range(1, 2 * n + 1):
            # 如果队首元素已经超出了当前考虑的子数组范围，则移除
            if dq and dq[0] < i - n:
                dq.popleft()

            # 计算当前子数组的和
            res = max(res, pre_sum[i] - pre_sum[dq[0]])

            # 维护单调队列
            while dq and pre_sum[dq[-1]] >= pre_sum[i]:
                dq.pop()

            dq.append(i)

        return res

    # 还要打印路径
    def maxSubArrayWithPath(self, nums: List[int]) -> Tuple[int, List[int]]:
        if not nums:
            return 0, []

        max_sum = current_max = nums[0]
        start = end = temp_start = 0

        for i in range(1, len(nums)):
            if current_max + nums[i] > nums[i]:
                current_max += nums[i]
            else:
                current_max = nums[i]
                temp_start = i

            if current_max > max_sum:
                max_sum = current_max
                start = temp_start
                end = i

        return max_sum, nums[start:end+1]
