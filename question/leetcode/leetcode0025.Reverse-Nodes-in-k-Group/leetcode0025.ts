class ListNode {
	val: number
	next: ListNode | null
	constructor(val?: number, next?: ListNode | null) {
		this.val = (val === undefined ? 0 : val)
		this.next = (next === undefined ? null : next)
	}

}

function reverseKGroupII(head: ListNode | null, k: number): ListNode | null {
	if (!head) {
		return null
	}
	let a: ListNode | null = head, b: ListNode | null = head
	for (let i = 0; i < k; i++) {
		if (!b) {
			return head
		}
		b = b.next

	}
	let newHead = reverseII(a, b)
	a.next = reverseKGroupII(b, k)
	return newHead;
};


function reverseII(head: ListNode | null, target: ListNode | null): ListNode | null {
	let pre: ListNode | null = null, cur: ListNode | null = head, next: ListNode | null = head;
	while (cur !== target) {
		if (!cur) {
			return pre
		}
		next = cur.next;
		cur.next = pre;
		pre = cur;
		cur = next;
	}
	return pre
}
