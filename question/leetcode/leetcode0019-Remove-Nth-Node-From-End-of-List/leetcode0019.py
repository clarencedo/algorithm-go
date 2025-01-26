class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def removeNthFromEnd(head, n):
    dummy = ListNode(0)
    dummy.next = head
    target = findNthFromEnd(dummy, n+1)
    target.next = target.next.next
    return dummy.next


def findNthFromEnd(head, n):
    if not head:
        return None
    fast = slow = head
    for i in range(n):
        if not fast:
            return None
        fast = fast.next
    while fast:
        fast = fast.next
        slow = slow.next
    return slow