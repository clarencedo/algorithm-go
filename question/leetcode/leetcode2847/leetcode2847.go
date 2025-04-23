package leetcode

import (
	"sort"
	"strconv"
)

// 给定一个正正数n，返回一个字符串，使得字符串中的每个字符都是数字并且字符串中的数字之乘积等于n
// 返回的字符串中的数字必须是最小的
// 如果这样的字符串不存在，返回"-1"
func smallestNumber(n int) string {
	if n < 10 {
		return strconv.Itoa(n)
	}

	var numbers []int
	for i := 9; i >= 2; i-- {
		if n%i == 0 {
			numbers = append(numbers, i)
			n /= i
		}
	}

	//不能被完全分解
	if n > 1 {
		return "-1"
	}
	sort.Ints(numbers)
	var ans string
	for _, v := range numbers {
		ans += strconv.Itoa(v)
	}
	return ans
}