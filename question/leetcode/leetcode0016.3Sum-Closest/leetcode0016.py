class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums.sort()
        n = len(nums)
        closest = nums[0] + nums[1] + nums[2]

        for i in range(n-2):
            left, right = i+1, n-1
            while left < right:
                cur_total = nums[i] + nums[left] + nums[right]
                if abs(target - cur_total) < abs(target - closest):
                    closest = cur_total
                if cur_total < target:
                    left += 1
                elif cur_total > target:
                    right -= 1
                else:
                    return target

        return closest
