class ListNode {
	val: number;
	next: ListNode | null;
	constructor(val?: number, next?: ListNode | null) {
		this.val = val === undefined ? 0 : val;
		this.next = next === undefined ? null : next;
	}
}

// 链表A长度为：a + c
// 链表B长度为：b + c
// c为公共长度部分
// 如果某个指针走到了链表末尾，就让它重新从链表头开始遍历
// 最后两者都走了a+b+c步，如果有交点，会在同一点相遇，否则都变成null
function getIntersectionNode(headA: ListNode | null, headB: ListNode | null): ListNode | null {
    let p1 = headA;
    let p2 = headB;
    while (p1 !== p2) {
        p1 = p1 ? p1.next : headB;
        p2 = p2 ? p2.next : headA;
    }
    return p1;
}
