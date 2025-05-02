package leetcode

import "clarencedu/algorithm-go/structure"

type TreeNode = structure.TreeNode

// 解法一：bfs
// 不需要维护整个中序遍历结果，只需要记录上一个访问的节点和当前节点，如果上一个访问的节点是p，那么当前节点就是后继节点
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var pre, cur *TreeNode = nil, root
	var queue []*TreeNode

	for len(queue) > 0 || cur != nil {
		for cur != nil {
			queue = append(queue, cur)
			cur = cur.Left
		}
		cur = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if pre == p {
			return cur
		}
		pre = cur
		cur = cur.Right
	}

	return nil
}

// 解法二：利用二叉搜索树的性质
func inorderSuccessorII(root *TreeNode, p *TreeNode) *TreeNode {
	var successor *TreeNode

	//如果p的右子树不为空，则p的中序后继一定在p的右子树中
	if p.Right != nil {
		successor = p.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		return successor
	}

	node := root
	//如果右子树为空，则中序后继一定在p的祖先节点中
	for node != nil {
		if node.Val > p.Val {
			//如果node节点的值大于p的值，则p的中序后继可能是node或者在node的左子树中,因此用node更新答案，并将node移动到其左子树中
			successor = node
			node = node.Left
		} else {
			//如果node节点的值小于等于p的值，则p的中序后继在node的右子树中，因此将node移动到其右子树中
			node = node.Right
		}
	}

	return successor
}