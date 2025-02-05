package leetcode

import "clarencedu/algorithm-go/structure"

type ListNode = structure.ListNode

// INFO: 寻找下一个更大的元素，经典的单调栈问题
// 和数组的单调栈问题类似，只不过这里是链表,外层循环遍历变成了移动链表指针遍历
func nextLargerNodes(head *ListNode) []int {
	var res []int
	var stack [][]int //存储元素值和索引,[[val,index],[val,index]]
	cur := head
	idx := -1

	for ; cur != nil; cur = cur.Next {
		idx++
		// 初始化res，占位，后面会被替换
		res = append(res, 0)
		// 找到比当前元素大的元素
		for len(stack) > 0 && stack[len(stack)-1][0] < cur.Val {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//这一题只需要保存值，不需要保存索引
			//否则 res[top[1]] = idx
			res[top[1]] = cur.Val
		}
		stack = append(stack, []int{cur.Val, idx})
	}

	return res
}