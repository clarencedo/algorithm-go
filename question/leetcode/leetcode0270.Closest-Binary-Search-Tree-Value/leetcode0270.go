package leetcode

import (
	"clarencedu/algorithm-go/structure"
	"math"
)

type TreeNode = structure.TreeNode

// 解法一：dfs
func closestValue(root *TreeNode, target float64) int {
	return findClosest(root, target, root.Val)
}

func findClosest(node *TreeNode, target float64, closest int) int {
	if node == nil {
		return closest
	}

	if math.Abs(float64(node.Val)-target) < math.Abs(float64(closest)-target) ||
		(math.Abs(float64(node.Val)-target) == math.Abs(float64(closest)-target) && node.Val < closest) {
		closest = node.Val
	}

	if target < float64(node.Val) {
		return findClosest(node.Left, target, closest)
	} else {
		return findClosest(node.Right, target, closest)
	}
}

// 解法二：dfs中序遍历,然后从有序结果中找最接近的值
func closestValueII(root *TreeNode, target float64) int {
	inorderRes := []int{}
	inorder(root, &inorderRes)
	return findClosestInorder(inorderRes, target)
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}

func findClosestInorder(arr []int, target float64) int {
	closestDiff := math.MaxFloat64
	closestVal := arr[0]
	for _, v := range arr {
		diff := math.Abs(float64(v) - target)
		if diff < closestDiff || (diff == closestDiff && v < closestVal) {
			closestDiff = diff
			closestVal = v
		}
	}
	return closestVal
}

// 解法三： 利用queue，中序遍历
func closestValueIII(root *TreeNode, target float64) int {
	closest := -1 << 62 // Use min int64 instead of min int / 2 for Go's int type
	minDifference := math.MaxFloat64
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		difference := math.Abs(float64(node.Val) - target)
		if difference < minDifference || (difference == minDifference && node.Val < closest) {
			closest = node.Val
			minDifference = difference
		}
		node = node.Right
	}
	return closest
}

// 解法四：Morris Traversal, 莫里斯遍历
func closestValueIV(root *TreeNode, target float64) int {
	closest := -1 << 62 // Minimum int64
	minDifference := math.MaxFloat64
	node := root
	for node != nil {
		if node.Left == nil {
			difference := math.Abs(float64(node.Val) - target)
			if difference < minDifference || (difference == minDifference && node.Val < closest) {
				closest = node.Val
				minDifference = difference
			}
			node = node.Right
		} else {
			predecessor := node.Left
			for predecessor.Right != nil && predecessor.Right != node {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				predecessor.Right = node
				node = node.Left
			} else {
				predecessor.Right = nil
				difference := math.Abs(float64(node.Val) - target)
				if difference < minDifference || (difference == minDifference && node.Val < closest) {
					closest = node.Val
					minDifference = difference
				}
				node = node.Right
			}
		}
	}
	return closest
}