from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        table = {}

        for i, v in enumerate(nums):
            complement = target - v
            if complement in table:
                return [table[complement], i]
            table[v] = i

        return []
