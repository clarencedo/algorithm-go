class Solution:
    def maxArea(self, height: List[int]) -> int:
        left, riight = 0, len(height) - 1
        capacity = 0
        while left < riight:
            cur_capacity = min(height[left], height[riight]) * (riight - left)
            if height[left] < height[riight]:
                left += 1
            else:
                riight -= 1
            if cur_capacity > capacity:
                capacity = cur_capacity
        return capacity
