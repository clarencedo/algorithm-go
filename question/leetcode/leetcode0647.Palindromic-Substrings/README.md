# [647. Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/)

## 题目

Given a string, your task is to count how many palindromic substrings in this string.

The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.

**Example 1:**

```
Input: "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".
```

**Example 2:**

```
Input: "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
```

**Note:**

1. The input string length won't exceed 1000.

## 题目大意

给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

## 解题思路

- 暴力解法，从左往右扫一遍字符串，以每个字符做轴，用中心扩散法，依次遍历计数回文子串。

## 代码

```go
package leetcode

func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += countPalindrome(s, i, i)
		res += countPalindrome(s, i, i+1)
	}
	return res
}

func countPalindrome(s string, left, right int) int {
	res := 0
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		left--
		right++
		res++
	}
	return res
}
```