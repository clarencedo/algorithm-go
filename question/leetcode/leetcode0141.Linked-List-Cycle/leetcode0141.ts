// class ListNode {
// 	val: number
// 	next: ListNode | null
// 	constructor(val?: number, next?: ListNode | null) {
// 		this.val = (val === undefined ? 0 : val)
// 		this.next = (next === undefined ? null : next)
// 	}
// }

function hasCycle(head: ListNode | null): boolean {
	if (head === null) {
		return false;
	}
	let slow = head, fast = head;
	let hasCircle: boolean = false;
	while (fast !== null && fast.next !== null) {
		slow = slow.next!;
		fast = fast.next.next!;
		if (slow === fast) {
			hasCircle = true;
			break;
		}
	}

	return hasCircle
};