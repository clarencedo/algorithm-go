package leetcode

import (
	"clarencedu/algorithm-go/structure"
	"fmt"
	"testing"
)

func Test_Problem382(t *testing.T) {
	header := structure.Ints2List([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	obj := Constructor(header)
	fmt.Printf("obj = %v\n", structure.List2Ints(header))
	param1 := obj.GetRandom()
	fmt.Printf("param_1 = %v\n", param1)
	param1 = obj.GetRandom()
	fmt.Printf("param_1 = %v\n", param1)
	param1 = obj.GetRandom()
	fmt.Printf("param_1 = %v\n", param1)
	param1 = obj.GetRandom()
	fmt.Printf("param_1 = %v\n", param1)
	param1 = obj.GetRandom()
	fmt.Printf("param_1 = %v\n", param1)
	param1 = obj.GetRandom()
	fmt.Printf("param_1 = %v\n", param1)

	fmt.Printf("\n\n\n")
}