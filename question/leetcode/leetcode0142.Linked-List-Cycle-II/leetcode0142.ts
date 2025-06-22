class ListNode {
	val: number
	next: ListNode | null
	constructor(val?: number, next?: ListNode | null) {
		this.val = (val === undefined ? 0 : val)
		this.next = (next === undefined ? null : next)
	}
}
function detectCycle2(head: ListNode | null): ListNode | null {
	if (head === null) return null;
	let fast = head;
	let slow = head;
	let hasCirle: boolean = false;
	while (fast !== null && fast.next !== null) {
		fast = fast.next.next!;
		slow = slow.next!;
		if (fast === slow) {
			hasCirle = true;
			break;
		}
	}
	if (!hasCirle) {
		return null;
	}
	// if (fast == null || fast.next ==null){
	//     return null;
	// }
	slow = head;
	while (slow != fast) {
		fast = fast.next!;
		slow = slow.next!;
	}
	return slow;
};
