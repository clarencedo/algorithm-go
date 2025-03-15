package template

import "clarencedu/algorithm-go/structure"

// func backtarck(state *State, choices []Choice, res *[]State) {
// 	if isSolution(state) {
// 		recordSolution(state, res)
// 		return
// 	}
// 	for _, choice := range choices {
// 		if isValid(state, choice) {
// 			makeChoice(state, choice)
// 			backtarck(state, choices, res)
// 			undochoice(state, choice)
// 		}
// 	}
// }

// 二叉树前序遍历 找值为7的节点, 若遇到3则提前返回（剪枝）
type TreeNode = structure.TreeNode

func preOrderII(root *TreeNode, res *[][]*TreeNode, path *[]*TreeNode) {
	if root == nil {
		return
	}
	// 尝试
	*path = append(*path, root)
	if root.Val == 7 {
		// 记录解
		*res = append(*res, append([]*TreeNode{}, *path...))
	}
	preOrderII(root.Left, res, path)
	preOrderII(root.Right, res, path)
	// 回退
	*path = (*path)[:len(*path)-1]
}

// 输入一个整数数组，数组中可能包含重复元素，返回所有不重复的排列。
func backtrackII(nums []int, path []int, used []bool, res *[][]int) {
	if len(path) == len(nums) {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backtrackII(nums, path, used, res)
		path = path[:len(path)-1]
		used[i] = false
	}
}