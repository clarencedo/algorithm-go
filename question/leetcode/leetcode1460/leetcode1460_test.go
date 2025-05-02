package leetcode

import "testing"

type question1460 struct {
	para1460
	ans1460
}

type para1460 struct {
	target []int
	arr    []int
}

type ans1460 struct {
	one bool
}

func TestCanBeEqual(t *testing.T) {
	qs := []question1460{
		{para1460{[]int{1, 2, 3, 4}, []int{2, 4, 1, 3}}, ans1460{true}},
		{para1460{[]int{7}, []int{7}}, ans1460{true}},
		{para1460{[]int{1, 12}, []int{12, 1}}, ans1460{true}},
		{para1460{[]int{3, 7, 9}, []int{3, 7, 11}}, ans1460{false}},
		{para1460{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}}, ans1460{true}},
	}

	for _, q := range qs {
		a, p := q.ans1460, q.para1460
		if got := canBeEqual(p.target, p.arr); got != a.one {
			t.Errorf("canBeEqual(%v, %v) = %v; want %v", p.target, p.arr, got, a.one)
		}
	}
}
