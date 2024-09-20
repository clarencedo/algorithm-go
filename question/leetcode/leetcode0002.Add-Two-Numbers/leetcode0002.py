# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        head = ListNode()
        carry = 0
        cur = head
        while l1 or l2 or carry:
            if l1 is not None:
                n1 = l1.val
                l1 = l1.next
            else:
                n1 = 0
            if l2 is not None:
                n2 = l2.val
                l2 = l2.next
            else:
                n2 = 0

            cur.next = ListNode((n1 + n2 + carry) % 10)
            cur = cur.next
            carry = (n1 + n2 + carry) // 10
        return head.next
