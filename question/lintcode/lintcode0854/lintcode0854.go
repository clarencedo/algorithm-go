package lintcode

import (
	"clarencedu/algorithm-go/structure"
	"math"
)

type TreeNode = structure.TreeNode

var (
	parentMap   map[*TreeNode]*TreeNode
	targetNode  *TreeNode
	minDepth    int
	closestLeaf *TreeNode
)

// FindClosestLeaf 找到距离目标节点最近的叶节点
func FindClosestLeaf(root *TreeNode, k int) int {
	parentMap = make(map[*TreeNode]*TreeNode)
	targetNode = nil
	minDepth = math.MaxInt32
	closestLeaf = nil

	findTarget(root, k)

	if targetNode == nil {
		return 0
	}

	findClosestLeafFromSubtree(targetNode, 0)

	findClosestLeafFromParent(targetNode)

	return closestLeaf.Val
}

func findTarget(root *TreeNode, k int) {
	if root == nil {
		return
	}
	if root.Val == k {
		targetNode = root
	}
	if root.Left != nil {
		parentMap[root.Left] = root
		findTarget(root.Left, k)
	}
	if root.Right != nil {
		parentMap[root.Right] = root
		findTarget(root.Right, k)
	}
}

func findClosestLeafFromSubtree(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	// 判断是否是叶子节点
	if root.Left == nil && root.Right == nil {
		if depth < minDepth {
			minDepth = depth
			closestLeaf = root
		}
		return
	}
	findClosestLeafFromSubtree(root.Left, depth+1)
	findClosestLeafFromSubtree(root.Right, depth+1)
}

func findClosestLeafFromParent(node *TreeNode) {
	visited := make(map[*TreeNode]bool)
	visited[node] = true
	distance := 0

	// 遍历父节点路径，查找最近的叶节点
	for current := node; current != nil; current = parentMap[current] {
		visited[current] = true
		// 查找从当前节点出发的叶节点
		findLeaf(current, distance, visited)
		distance++
	}
}

func findLeaf(node *TreeNode, distance int, visited map[*TreeNode]bool) {
	if node == nil || visited[node] {
		return
	}
	// 判断是否是叶子节点
	if node.Left == nil && node.Right == nil {
		if distance < minDepth {
			minDepth = distance
			closestLeaf = node
		}
		return
	}
	// 继续递归搜索左右子树
	findLeaf(node.Left, distance+1, visited)
	findLeaf(node.Right, distance+1, visited)
}