from collections import defaultdict


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        mp = {}
        for s in strs:
            str = "".join(sorted(s))
            if str not in mp:
                mp[str] = [s]
            else:
                mp[str].append(s)

        ans = []
        for k, v in mp.items():
            ans.append(v)

        return ans

    # 如果输出要求字典序
    # 比如输入 ["eat", "tea", "tan", "ate", "nat", "bat"]
    # 输出 [["ate", "eat", "tea"], ["nat", "tan"], ["bat"]]
    def groupAnagramsII(self, strs: List[str]) -> List[List[str]]:
        mp = defaultdict(list)
        for s in strs:
            sorted_str = "".join(sorted(s))
            mp[sorted_str].append(s)

        # 每组内部排序
        groups = [sorted(group) for group in mp.values()]

        # 每组第一个单词进行排序
        groups.sort(key=lambda x: x[0])

        return groups
