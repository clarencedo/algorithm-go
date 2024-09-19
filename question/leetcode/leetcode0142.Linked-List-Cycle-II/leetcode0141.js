var detectCycle = function (head) {
  let fast = head;
  let slow = head;
  let hasCirle = false;
  while (fast !== null && fast.next !== null) {
    fast = fast.next.next;
    slow = slow.next;
    if (fast === slow) {
      hasCirle = true;
      break;
    }
  }
  if (!hasCirle) {
    return null;
  }
  slow = head;
  while (slow != fast) {
    fast = fast.next;
    slow = slow.next;
  }
  return slow;
};