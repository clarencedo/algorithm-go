class Solution:
    def search(self, nums: List[int], target: int) -> int:
        n = len(nums)
        if n == 0:
            return -1
        if n == 1 and nums[0] == target:
            return 0

        left, right = 0, n - 1
        while left <= right:
            mid = left + (right - left) // 2
            if nums[mid] == target:
                return mid
            elif nums[0] <= nums[mid]:
                # 左边有序
                if nums[left] <= target < nums[mid]:
                    # target在左边
                    right = mid - 1
                else:
                    left = mid + 1
            else:
                # 右边有序
                if nums[mid] < target <= nums[right]:
                    # target在右边
                    left = mid + 1
                else:
                    right = mid - 1
        return -1
