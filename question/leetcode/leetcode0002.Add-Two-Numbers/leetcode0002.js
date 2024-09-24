var addTwoNumbers = function (l1, l2) {
  let dummy = new ListNode(0);
  let n1,
    n2,
    carry = 0,
    current = dummy;
  while (l1 !== null || l2 !== null || carry !== 0) {
    if (l1 === null) {
      n1 = 0;
    } else {
      n1 = l1.val;
      l1 = l1.next;
    }
    if (l2 === null) {
      n2 = 0;
    } else {
      n2 = l2.val;
      l2 = l2.next;
    }

    current.next = new ListNode((n1 + n2 + carry) % 10);
    current = current.next;
    carry = Math.floor((n1 + n2 + carry) / 10);
  }
  return dummy.next;
};
