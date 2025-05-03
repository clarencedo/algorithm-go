class ListNode {
    int val;
    ListNode next;

    ListNode() {
    }

    public ListNode(int val) {
        this.val = val;
    }

    public ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }
}

class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode dummy = new ListNode(0);
        Integer n1, n2, carry = 0;
        var cur = dummy;
        while (l1 != null || l2 != null || carry != 0) {
            if (l1 == null) {
                n1 = 0;
            } else {
                n1 = l1.val;
                l1 = l1.next;
            }
            if (l2 == null) {
                n2 = 0;
            } else {
                n2 = l2.val;
                l2 = l2.next;
            }

            cur.next = new ListNode((n1 + n2 + carry) % 10);
            cur = cur.next;
            carry = (n1 + n2 + carry) / 10;

        }
        return dummy.next;

    }
}
