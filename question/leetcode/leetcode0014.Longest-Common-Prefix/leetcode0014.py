class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        prefix = ""
        firstStr = strs[0]
        for i in range(len(firstStr)):
            char = firstStr[i]
            for j in range(1, len(strs)):
                if i >= len(strs[j]) or strs[j][i] != char:
                    return prefix
            prefix += char
        
        return prefix
    