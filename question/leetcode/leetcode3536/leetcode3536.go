package leetcode

import (
	"sort"
	"strconv"
)

// 排序
func maxProduct(n int) int {
	// 转成字符串
	s := strconv.Itoa(n)

	// 转成 rune 数组或 byte 数组进行排序
	digits := []byte(s)
	sort.Slice(digits, func(i, j int) bool {
		return digits[i] < digits[j]
	})

	// 拿出最后两个字符，转成数字后相乘
	d1 := int(digits[len(digits)-1] - '0')
	d2 := int(digits[len(digits)-2] - '0')
	return d1 * d2
}

// 不排序
func maxProductII(n int) int {
	max1, max2 := 0, 0

	for n > 0 {
		d := n % 10
		if d > max1 {
			max1, max2 = d, max1
		} else if d > max2 {
			max2 = d
		}
		n /= 10
	}

	return max1 * max2
}
