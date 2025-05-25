from typing import List


class Solution:
    def widestPairOfIndices(self, nums1: List[int], nums2: List[int]) -> int:
        '''
        找到最宽的下标对(i,j)，使得nums1[i]到nums1[j]的和等于nums2[i]到nums2[j]的和
        即：nums1[i] + ... + nums1[j] == nums2[i] + ... + nums2[j]
        转化为：(nums1[i] - nums2[i]) + ... + (nums1[j] - nums2[j]) == 0
        使用前缀和差值和哈希表来解决
        '''
        # 初始化结果和哈希表
        max_width = 0
        prefix_diff = 0  # 前缀和差值
        diff_to_index = {0: -1}  # 差值到索引的映射，初始化为{0:-1}表示空前缀,记录每个前缀和差值 第一次出现的位置

        # 遍历数组
        for i in range(len(nums1)):
            # 计算当前位置的前缀和差值
            prefix_diff += nums1[i] - nums2[i]

            # 如果这个差值之前出现过，说明找到了一个有效的下标对
            # 如果之前prefix_diff_i = x
            # 当前prefix_diff_j = x
            # 那么nums1[i] + ... + nums1[j] == nums2[i] + ... + nums2[j]
            if prefix_diff in diff_to_index:
                max_width = max(max_width, i - diff_to_index[prefix_diff])
            else:
                # 如果是第一次出现这个差值，记录它的索引
                diff_to_index[prefix_diff] = i

        return max_width
