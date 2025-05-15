class Solution:

    # 不跨越数组边界：即普通的最大子数组和问题（Kadane算法可解决）
    # 跨越数组边界：这种情况可以转化为"总和减去最小子数组和"
    def maxSubarraySumCircular(self, nums: List[int]) -> int:
        total = sum(nums)
        cur_max = max_sum = nums[0]
        cur_min = min_sum = nums[0]

        for num in nums[1:]:
            cur_max = max(num, cur_max + num)
            max_sum = max(max_sum, cur_max)

            cur_min = min(num, cur_min + num)
            min_sum = min(min_sum, cur_min)

        if max_sum < 0:
            return max_sum
        return max(max_sum, total - min_sum)
