package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums))
}

func build(nums []int, left, right int) *TreeNode {
	if left >= right {
		return nil
	}

	maxIdx := left
	for i := left; i < right; i++ {
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}

	root := &TreeNode{Val: nums[maxIdx]}
	root.Left = build(nums, left, maxIdx)
	root.Right = build(nums, maxIdx+1, right)
	return root
}