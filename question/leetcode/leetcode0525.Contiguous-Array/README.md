# [525. Contiguous Array](https://leetcode.com/problems/contiguous-array/)


## 题目

Given a binary array `nums`, return *the maximum length of a contiguous subarray with an equal number of* `0` *and* `1`.

**Example 1:**

```
Input: nums = [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.
```

**Example 2:**

```
Input: nums = [0,1,0]
Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
```

**Constraints:**

- `1 <= nums.length <= 105`
- `nums[i]` is either `0` or `1`.

## 题目大意

给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。

## 解题思路

- 0 和 1 的数量相同可以转化为两者数量相差为 0，如果将 0 看作为 -1，那么原题转化为求最长连续子数组，其元素和为 0 。又变成了区间内求和的问题，自然而然转换为前缀和来处理。假设连续子数组是 [i,j] 区间，这个区间内元素和为 0 意味着 prefixSum[j] - prefixSum[i] = 0，也就是 prefixSum[i] = prefixSum[j]。不断累加前缀和，将每个前缀和存入 map 中。一旦某个 key 存在了，代表之前某个下标的前缀和和当前下标构成的区间，这段区间内的元素和为 0 。这个区间是所求。扫完整个数组，扫描过程中动态更新最大区间长度，扫描完成便可得到最大区间长度，即最长连续子数组。

## 代码

```go
package leetcode

func findMaxLength(nums []int) int {
	dict := map[int]int{}
	dict[0] = -1
	count, res := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count--
		} else {
			count++
		}
		if idx, ok := dict[count]; ok {
			res = max(res, i-idx)
		} else {
			dict[count] = i
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
