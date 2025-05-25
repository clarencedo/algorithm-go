class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        i = n - 2

        # Step 1: Find the first decreasing element
        while i >= 0 and nums[i] >= nums[i + 1]:
            i -= 1

        # If no such element, reverse the whole array
        if i < 0:
            nums.reverse()
            return

        # Step 2: Find the smallest number > nums[i] in the right part
        j = n - 1
        while j >= 0 and nums[j] <= nums[i]:
            j -= 1

        # Step 3: Swap nums[i] and nums[j]
        nums[i], nums[j] = nums[j], nums[i]

        # Step 4: Reverse the part after i
        left, right = i + 1, n - 1
        while left < right:
            nums[left], nums[right] = nums[right], nums[left]
            left += 1
            right -= 1
