class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        window = {}
        left, right = 0, 0
        res = 0
        while right < len(s):
            rightChar = s[right]
            window[rightChar] = window.get(rightChar, 0) + 1
            right += 1
            while window[rightChar] > 1:
                leftChar = s[left]
                window[leftChar] = window.get(leftChar, 0) - 1
                left += 1
            res = max(res, right - left)
        return res
