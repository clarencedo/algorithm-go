
class Solution:
    def longestPalindrome(self, s: str) -> str:
        if len(s) < 2:
            return s

        windowSize = len(s)

        for i in range(windowSize, 0, -1):
            for j in range(0, windowSize - i + 1):
                substr = s[j:j+i]
                if substr == substr[::-1]:
                    return substr

        return s[0]
