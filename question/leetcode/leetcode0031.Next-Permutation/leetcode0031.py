from typing import List


class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        i = n - 2

        while i >= 0 and nums[i] >= nums[i + 1]:
            i -= 1

        if i < 0:
            nums.reverse()
            return

        j = n - 1
        while j >= 0 and nums[j] <= nums[i]:
            j -= 1

        nums[i], nums[j] = nums[j], nums[i]

        # 手动翻转
        # left, right = i + 1, n - 1
        # while left < right:
        #     nums[left], nums[right] = nums[right], nums[left]
        #     left += 1
        #     right -= 1

        # nums[i+1:] = nums[i+1:][::-1]

        nums[i + 1 :] = reversed(nums[i + 1 :])