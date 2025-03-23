# [78. Subsets](https://leetcode.com/problems/subsets/)

## 题目

Given a set of **distinct** integers, *nums*, return all possible subsets (the power set).

**Note:** The solution set must not contain duplicate subsets.

**Example:**

    Input: nums = [1,2,3]
    Output:
    [
      [3],
      [1],
      [2],
      [1,2,3],
      [1,3],
      [2,3],
      [1,2],
      []
    ]

## 题目大意

给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。说明：解集不能包含重复的子集。

## 解题思路

- 找出一个集合中的所有子集，空集也算是子集。且数组中的数字不会出现重复。用 DFS 暴力枚举即可。
- 这一题和第 90 题，第 491 题类似，可以一起解答和复习。

假设 nums = [1,2,3]，那么回溯的过程如下：

```
index=0, current=[]       → 加入结果 [[]]
  ├── index=1, current=[1] → 加入结果 [[], [1]]
  │    ├── index=2, current=[1,2] → 加入结果 [[], [1], [1,2]]
  │    │    ├── index=3, current=[1,2,3] → 加入结果 [[], [1], [1,2], [1,2,3]]
  │    │    └── 回溯，current=[1,2]
  │    ├── index=3, current=[1,3] → 加入结果 [[], [1], [1,2], [1,2,3], [1,3]]
  │    └── 回溯，current=[1]
  ├── index=2, current=[2] → 加入结果 [[], [1], [1,2], [1,2,3], [1,3], [2]]
  │    ├── index=3, current=[2,3] → 加入结果 [[], [1], [1,2], [1,2,3], [1,3], [2], [2,3]]
  │    └── 回溯，current=[2]
  ├── index=3, current=[3] → 加入结果 [[], [1], [1,2], [1,2,3], [1,3], [2], [2,3], [3]]
  └── 回溯，current=[]
```

最终得到：

```
[[], [1], [1,2], [1,2,3], [1,3], [2], [2,3], [3]]
```