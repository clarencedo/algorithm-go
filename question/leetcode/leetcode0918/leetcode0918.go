package leetcode

// TODO:
// 方法一：动态规划
func maxSubarraySumCircular(nums []int) int {
	dp := make(chan int, len(nums))
	return -1
}

// 方法二：前缀和+ 单调队列，把环形数组转换成两个数组
// 问题转化成了 在长度为2n的数组上找一个长度为n的连续子数组，使得和最大
// 前缀和数组preSum[i]表示nums[0]到nums[i-1]的和
// 目标区间就是pre[j]-pre[i]最大，即pre[i]最小
// 单调递增队列 就是方便找到这个pre[i]最小值，队列最后一个元素就是pre[i]最大值
//
// 这个解法的核心就是
// 找出前缀和pre[j]-pre[i]的最大值，单调队列是帮助找到这个最大差值的方法(队头和队尾的差值最大)
// 队列维护的元素是[index,前缀和的值]这个结构
func maxSubarraySumCircularII(nums []int) int {
	type pair struct{ idx, val int }
	n := len(nums)
	pre, res := nums[0], nums[0]
	q := []pair{{0, pre}}
	for i := 1; i < 2*n; i++ {
		//维护一个长度为n的窗口,计算pre[i]只考虑长度最多为n的子数组,如果长度超出n,则移除最左边的元素
		for len(q) > 0 && q[0].idx < i-n {
			q = q[1:]
		}
		pre += nums[i%n]
		res = max(res, pre-q[0].val)
		// 维护一个单调递增队列,如果当前前缀和比队列最后一个元素大,则移除队列最后一个元素,保持队列的单调递增性
		for len(q) > 0 && q[len(q)-1].val >= pre {
			q = q[:len(q)-1]
		}
		q = append(q, pair{i, pre})
	}

	return res
}
