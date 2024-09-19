def detect_cycle(head):
    if not head or not head.next:  # 如果链表为空或只有一个节点，则没有环
        return None

    fast = slow = head
    has_circle = False
    while fast and fast.next:
        fast = fast.next.next
        slow = slow.next
        if fast == slow:
            has_circle = True
            break
    if not has_circle:
        return None
    slow = head
    while slow != fast:
        fast = fast.next
        slow = slow.next
    return slow