class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        sorted_nums = [num * num for num in nums]
        sorted_nums.sort()
        return sorted_nums

    def sortedSquaresII(self, nums: List[int]) -> List[int]:
        n = len(nums)
        ans = [0] * n
        left, right = 0, n - 1
        while left <= right:
            if nums[left] * nums[left] > nums[right] * nums[right]:
                ans[n - 1] = nums[left] * nums[left]
                left += 1
            else:
                ans[n - 1] = nums[right] * nums[right]
                right -= 1
            n -= 1
        return ans
