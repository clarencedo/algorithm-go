class ListNode {
	val: number;
	next: ListNode | null;
	constructor(val?: number, next?: ListNode | null) {
		this.val = val === undefined ? 0 : val;
		this.next = next === undefined ? null : next;
	}
}

function isPalindrome(head: ListNode | null): boolean {
	let vals: number[] = [];
	while (head) {
		vals.push(head.val);
		head = head.next;
	}
	const len = vals.length

	for (let i = 0; i < len / 2; i++) {
		if (vals[i] != vals[len - 1 - i]) {
			return false;
		}
	}
	return true;
}