package leetcode

import (
	"clarencedu/algorithm-go/structure"
	"fmt"
	"testing"
)

type question138 struct {
	para138
	ans138
}

// para 是参数
// one 代表第一个参数
type para138 struct {
	head [][]int
}

// ans 是答案
// one 代表第一个答案
type ans138 struct {
	one [][]int
}

func Test_Problem138(t *testing.T) {

	qs := []question138{

		{
			para138{[][]int{{7, structure.NULL}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}},
			ans138{[][]int{{7, structure.NULL}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}},
		},

		{
			para138{[][]int{{1, 1}, {2, 1}}},
			ans138{[][]int{{1, 1}, {2, 1}}},
		},

		{
			para138{[][]int{{3, structure.NULL}, {3, 0}, {3, structure.NULL}}},
			ans138{[][]int{{3, structure.NULL}, {3, 0}, {3, structure.NULL}}},
		},

		{
			para138{[][]int{}},
			ans138{[][]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 138------------------------\n")

	for _, q := range qs {
		_, p := q.ans138, q.para138
		fmt.Printf("【input】:%v       【output】:%v\n", p, node2Ints(copyRandomList(ints2Node(p.head))))
	}
	fmt.Printf("\n\n\n")
}

func ints2Node(one [][]int) *Node {
	nodesMap := map[int]*Node{}
	for index, nodes := range one {
		tmp := &Node{Val: nodes[0]}
		nodesMap[index] = tmp
	}
	for index, node := range nodesMap {
		if index != len(one)-1 {
			node.Next = nodesMap[index+1]
		} else {
			node.Next = nil
		}
	}
	for index, nodes := range one {
		if nodes[1] == structure.NULL {
			nodesMap[index].Random = nil
		} else {
			nodesMap[index].Random = nodesMap[nodes[1]]
		}
	}
	return nodesMap[0]
}

func node2Ints(head *Node) (res [][]int) {
	nodesMap, count, cur := map[*Node]int{}, 0, head
	if head == nil {
		return [][]int{}
	}
	for cur.Next != nil {
		nodesMap[cur] = count
		count++
		cur = cur.Next
	}
	nodesMap[cur] = count
	tmp := []int{}
	cur = head
	for cur.Next != nil {
		tmp := append(tmp, cur.Val)
		if cur.Random == nil {
			tmp = append(tmp, structure.NULL)
		} else {
			tmp = append(tmp, nodesMap[cur.Random])
		}
		res = append(res, tmp)
		cur = cur.Next
	}
	tmp = append(tmp, cur.Val)
	if cur.Random == nil {
		tmp = append(tmp, structure.NULL)
	} else {
		tmp = append(tmp, nodesMap[cur.Random])
	}
	res = append(res, tmp)
	return res
}