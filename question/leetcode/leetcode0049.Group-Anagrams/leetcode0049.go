package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
	//key是字符串，值是字符串切片
	//key用于存储排序后的字符串,值则存储原始字符串列表
	mp := map[string][]string{}
	for _, str := range strs {
		//字符串是不可变的,不能直接排序,转成byte slice排序
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}