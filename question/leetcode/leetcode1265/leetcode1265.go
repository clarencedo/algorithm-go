package leetcode

/*   Below is the interface for ImmutableListNode, which is already defined for you.
 *
 *   type ImmutableListNode struct {
 *
 *   }
 *
 *   func (this *ImmutableListNode) getNext() ImmutableListNode {
 *		// return the next node.
 *   }
 *
 *   func (this *ImmutableListNode) printValue() {
 *		// print the value of this node.
 *   }
 */

// TODO: not implemented
type ImmutableListNode struct {
	Val int
}

func (this *ImmutableListNode) getNext() ImmutableListNode {
	return ImmutableListNode{}
}

func (this *ImmutableListNode) printValue() {
}

// func printLinkedListInReverse(head ImmutableListNode) {
// 	if head != nil {
// 		printLinkedListInReverse(head.getNext())
// 		head.printValue()
// 	}
// }
