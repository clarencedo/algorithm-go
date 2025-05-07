package leetcode

import "clarencedu/algorithm-go/structure"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structure.TreeNode

type CBTInserter struct {
	root *TreeNode
	//队列，保存还没填满左右孩子的节点
	q []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	q := []*TreeNode{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left == nil || node.Right == nil {
			q = append(q, node)
		}

	}
	return CBTInserter{root, q}
}

func (this *CBTInserter) Insert(val int) int {
	parent := this.q[0]
	newNode := &TreeNode{Val: val}
	if parent.Left == nil {
		parent.Left = newNode
	} else {
		parent.Right = newNode
		//父节点已经满了，从队列移除
		this.q = this.q[1:]
	}
	this.q = append(this.q, newNode)
	return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
