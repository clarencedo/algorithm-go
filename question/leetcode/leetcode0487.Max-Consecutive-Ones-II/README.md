# [487. 最大连续1的个数](https://leetcode.cn/problems/max-consecutive-ones-ii/?envType=study-plan-v2&envId=premium-algo-100)

## 题目

Given a binary array, find the maximum number of consecutive 1s in this array if you can flip at most one 0.

Example 1:
Input: [1,0,1,1,0]
Output: 4
Explanation: Flip the first zero will get the the maximum number of consecutive 1s.
After flipping, the maximum number of consecutive 1s is 4.
Note:

The input array will only contain 0 and 1.
The length of input array is a positive integer and will not exceed 10,000
Follow up:
What if the input numbers come in one by one as an infinite stream? In other words, you can't store all numbers coming from the stream as it's too large to hold in memory. Could you solve it efficiently?

## 题目大意

给定一个二进制数组 nums ，如果最多可以翻转一个 0 ，则返回数组中连续 1 的最大个数。

## 解题思路

滑动窗口，维护一个窗口，窗口内最多只有一个 0，窗口的大小就是最大连续 1 的个数。