from typing import List


def productExceptSelf(nums: List[int]) -> List[int]:
    n = len(nums)
    ans = [1] * n

    # 左边所有数的乘积
    for i in range(1, n):
        ans[i] = ans[i - 1] * nums[i - 1]

    # 右边所有数的乘积
    right = 1
    for i in range(n - 2, -1, -1):
        right *= nums[i + 1]
        ans[i] *= right

    return ans
