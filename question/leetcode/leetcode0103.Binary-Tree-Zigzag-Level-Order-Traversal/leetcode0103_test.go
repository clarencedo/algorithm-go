package leetcode

import (
	"clarencedu/algorithm-go/structure"
	"fmt"
	"testing"
)

type question103 struct {
	para103
	ans103
}

// para 是参数
// one 代表第一个参数
type para103 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans103 struct {
	one [][]int
}

func Test_Problem103(t *testing.T) {

	qs := []question103{

		{
			para103{[]int{}},
			ans103{[][]int{}},
		},

		{
			para103{[]int{1}},
			ans103{[][]int{{1}}},
		},

		{
			para103{[]int{3, 9, 20, structure.NULL, structure.NULL, 15, 7}},
			ans103{[][]int{{3}, {9, 20}, {15, 7}}},
		},

		{
			para103{[]int{1, 2, 3, 4, structure.NULL, structure.NULL, 5}},
			ans103{[][]int{{1}, {3, 2}, {4, 5}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 103------------------------\n")

	for _, q := range qs {
		_, p := q.ans103, q.para103
		fmt.Printf("【input】:%v      ", p)
		root := structure.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", zigzagLevelOrder(root))
	}
	fmt.Printf("\n\n\n")
}