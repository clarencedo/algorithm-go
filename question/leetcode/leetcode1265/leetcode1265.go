package leetcode

// TODO: not implemented
type ImmutableListNode struct {
	Val int
}

func (this *ImmutableListNode) getNext() ImmutableListNode {
	return ImmutableListNode{}
}

func (this *ImmutableListNode) printValue() {
}

func printLinkedListInReverse(head ImmutableListNode) {
	if head != nil {
		printLinkedListInReverse(head.getNext())
		head.printValue()
	}
}
