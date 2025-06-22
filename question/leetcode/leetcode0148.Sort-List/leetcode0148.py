from typing import List


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

    def mergeSort(head: ListNode, tail: ListNode) -> ListNode:
        if not head:
            return head
        if head.next == tail:
            head.next = None
            return head

        slow, fast = head, head
        while fast != tail:
            slow = slow.next
            fast = fast.next
            if fast and fast.next != tail:
                fast = fast.next

        mid = slow
        return mergeTwoLists(mergeSort(head, mid), mergeSort(mid, tail))

    def mergeTwoLists(l1: ListNode, l2: ListNode) -> ListNode:
        dummy = ListNode()
        current = dummy

        while l1 and l2:
            if l1.val < l2.val:
                current.next, l1 = l1, l1.next
            else:
                current.next, l2 = l2, l2.next
            current = current.next  # 别忘了移动 current 指针

        current.next = l1 if l1 else l2  # 直接连接剩余部分

        return dummy.next

    def sortList(head: ListNode) -> ListNode:
        return mergeSort(head, None)
