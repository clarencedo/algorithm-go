package leetcode

import (
    "testing"
    "clarencedu/algorithm-go/structure"
)

func TestCBTInserter(t *testing.T) {
    root := &structure.TreeNode{Val: 1}
    inserter := Constructor(root)
    inserter.Insert(2)
    inserter.Insert(3)
    inserter.Insert(4)
    inserter.Insert(5)

    if inserter.Get_root().Val != 1 {
        t.Errorf("Expected root value to be 1, got %d", inserter.Get_root().Val)
    }

    if inserter.Get_root().Left.Val != 2 {
        t.Errorf("Expected left child of root to be 2, got %d", inserter.Get_root().Left.Val)
    }

    if inserter.Get_root().Right.Val != 3 {
        t.Errorf("Expected right child of root to be 3, got %d", inserter.Get_root().Right.Val)
    }
}
