class ListNode {
	val: number
	next: ListNode | null
	constructor(val?: number, next?: ListNode | null) {
		this.val = (val === undefined ? 0 : val)
		this.next = (next === undefined ? null : next)
	}
}

function sortList(head: ListNode | null): ListNode | null {
	return mergeSort(head, null);
};


// 从head开始，到Tail之前结束
const mergeSort = (head: ListNode | null, tail: ListNode | null) => {
	if (head == null) {
		return null
	}
	if (head.next == tail) {
		head.next = null
		return head;
	}
	let slow: ListNode = head, fast: ListNode = head;
	while (fast !== tail && fast.next !== tail) {
		slow = slow.next!;
		fast = fast.next?.next!;
	}
	let mid = slow;
	return mergeTwoList(mergeSort(head, mid), mergeSort(mid, tail))
}

const mergeTwoList = (l1: ListNode | null, l2: ListNode | null) => {
	let dummy: ListNode = new ListNode();
	let current = dummy;

	while (l1 != null && l2 != null) {
		if (l1.val < l2.val) {
			current.next = l1;
			l1 = l1.next;
		} else {
			current.next = l2;
			l2 = l2.next;
		}
		current = current.next;
	}

	if (l1 != null) {
		current.next = l1;
	}
	if (l2 != null) {
		current.next = l2;
	}

	return dummy.next;
}