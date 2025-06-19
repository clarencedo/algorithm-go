class Solution:

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
