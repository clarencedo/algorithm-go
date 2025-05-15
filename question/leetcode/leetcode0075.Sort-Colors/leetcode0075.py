class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.

        荷兰国旗问题
        三指针，left指向0的最右边界，right指向2的最左边界
        遍历数组，如果遇到0则与left交换，left右移；如果遇到2则与right交换，right左移
        时间复杂度O(n)，空间复杂度O(1)
        """
        left, right = 0, len(nums) - 1
        i = 0
        while i <= right:
            if nums[i] == 0:
                nums[i], nums[left] = nums[left], nums[i]
                left += 1
                i += 1
            elif nums[i] == 2:
                nums[i], nums[right] = nums[right], nums[i]
                right -= 1
                # 注意这里不增加i，因为交换后的新元素还需要检查
            else:  # nums[i] == 1
                i += 1
