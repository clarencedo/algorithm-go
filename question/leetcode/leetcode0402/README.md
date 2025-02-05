# [402. Remove K Digits](https://leetcode.com/problems/remove-k-digits/)

# [402. 移掉K位数字](https://leetcode-cn.com/problems/remove-k-digits/)

## 题目

Given a non-negative integer num represented as a string, remove k digits from the number so that the new number is the smallest possible.

Note:

- The length of num is less than 10002 and will be ≥ k.
- The given num does not contain any leading zero.

Example 1:

```c
Input: num = "1432219", k = 3
Output: "1219"
Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
```

Example 2:

```c
Input: num = "10200", k = 1
Output: "200"
Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.
```

Example 3:

```c
Input: num = "10", k = 2
Output: "0"
Explanation: Remove all the digits from the number and it is left with nothing which is 0.
```

## 题目大意

给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。

## 解题思路

典型的单调栈问题。维护一个单调递增的栈，当遇到比栈顶元素小的元素时，出栈，直到栈为空或者栈顶元素小于当前元素。最后将栈中元素拼接起来即可。注意最后的结果可能会有前导0，需要去掉。
