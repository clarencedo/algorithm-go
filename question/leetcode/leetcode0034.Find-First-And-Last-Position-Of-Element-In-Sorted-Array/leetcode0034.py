class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        left = self.findLeftBound(nums, target)
        right = self.findRightBound(nums, target)
        return [left, right]

    def findLeftBound(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = left + (right - left) // 2
            if nums[mid] < target:
                left = mid + 1
            elif nums[mid] > target:
                right = mid - 1
            elif nums[mid] == target:
                # 缩减右边界,继续往左边找,锁定左侧边界
                right = mid - 1
        # 循环结束时，left指向第一个大于target的元素，right指向最后一个小于等于target的元素
        # 检查left是否越界，可能所有元素都小于target
        # nums[2,4,6] ,target =5
        # 循环结束, left = 2 ,right = 1
        if left < len(nums) and nums[left] == target:
            return left
        return -1

    def findRightBound(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = left + (right - left) // 2
            if nums[mid] < target:
                left = mid+1
            elif nums[mid] > target:
                right = mid - 1
            elif nums[mid] == target:
                # 缩减左边界,继续往右边找,锁定右侧边界
                left = mid + 1

        if right >= 0 and nums[right] == target:
            return right

        return -1
