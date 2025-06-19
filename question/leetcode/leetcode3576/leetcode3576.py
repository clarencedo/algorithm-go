class Solution:
    def canMakeEqual(self, nums: List[int], k: int) -> bool:
        def f(target):
            arr = nums[:]  # 不修改原数组
            n = len(arr)
            count = 0
            for i in range(n - 1):
                if arr[i] != target:
                    # 操作：翻转当前和下一个
                    arr[i] *= -1
                    arr[i + 1] *= -1
                    count += 1
                    if count > k:
                        return False
            return all(x == target for x in arr)
        
        return f(1) or f(-1)