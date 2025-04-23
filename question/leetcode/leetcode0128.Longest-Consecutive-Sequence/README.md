# [128. Longest Consecutive Sequence](https://leetcode.com/problems/longest-consecutive-sequence/)

## 题目

Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

Your algorithm should run in O(_n_) complexity.

**Example:**

    Input: [100, 4, 200, 1, 3, 2]
    Output: 4
    Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

## 题目大意

给定一个未排序的整数数组，找出最长连续序列的长度。要求算法的时间复杂度为 O(n)。

## 解题思路

- 给出一个数组，要求找出最长连续序列，输出这个最长的长度。要求时间复杂度为 `O(n)`。
- 这一题可以先用暴力解决解决，代码见解法三。思路是把每个数都存在 `map` 中，先删去 `map` 中没有前一个数 `nums[i]-1` 也没有后一个数 `nums[i]+1` 的数 `nums[i]`，这种数前后都不连续。然后在 `map` 中找到前一个数 `nums[i]-1` 不存在，但是后一个数 `nums[i]+1` 存在的数，这种数是连续序列的起点，那么不断的往后搜，直到序列“断”了。最后输出最长序列的长度。
