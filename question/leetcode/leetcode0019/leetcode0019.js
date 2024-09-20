function removeNthFromEnd(head, n) {
  let dummy = new ListNode(-1);
  dummy.next = head;
  let target = findNthFromEnd(dummy, n + 1);
  target.next = target.next.next;
  return dummy.next;
}

function findNthFromEnd(head, n) {
  let fast = head;
  for (let i = 0; i < n; i++) {
    fast = fast.next;
  }
  let slow = head;
  while (fast !== null) {
    fast = fast.next;
    slow = slow.next;
  }
  return slow;
}