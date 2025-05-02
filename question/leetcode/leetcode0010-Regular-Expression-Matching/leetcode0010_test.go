package leetcode

import (
	"testing"
)

type question10 struct {
	para10
	ans10
}

type para10 struct {
	s string
	p string
}

type ans10 struct {
	one bool
}

func Test_Problem10(t *testing.T) {
	qs := []question10{
		{para10{"aa", "a"}, ans10{false}},
		{para10{"aa", "a*"}, ans10{true}},
		{para10{"ab", ".*"}, ans10{true}},
		{para10{"aab", "c*a*b"}, ans10{true}},
		{para10{"mississippi", "mis*is*p*."}, ans10{false}},
	}

	for _, q := range qs {
		a, p := q.ans10, q.para10
		if got := isMatch(p.s, p.p); got != a.one {
			t.Errorf("isMatch(%v, %v) = %v; want %v", p.s, p.p, got, a.one)
		}
	}
}
