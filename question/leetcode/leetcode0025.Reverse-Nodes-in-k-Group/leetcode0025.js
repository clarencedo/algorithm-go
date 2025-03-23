/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} k
 * @return {ListNode}
 */

const reverseKGroup = function (head, k) {
	if (!head) {
		return null
	}
	let a = head, b = head;
	for (let i = 0; i < k; i++) {
		if (!b) {
			return head
		}
		b = b.next
	}
	let newHead = reverse(a, b)
	a.next = reverseKGroup(b, k)

	return newHead
};

const reverse = function (head, target) {
	let pre, cur, next;
	pre = null, cur = head; next = head;
	while (cur !== target) {
		next = cur.next;
		cur.next = pre;
		pre = cur;
		cur = next;
	}

	return pre;
};
