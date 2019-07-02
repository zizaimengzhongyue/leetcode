class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        if l1 is None:
            return l2
        if l2 is None:
            return l1
        if l1.val < l2.val:
            l3 = l1
            l1 = l1.next
        else:
            l3 = l2
            l2 = l2.next
        head = l3
        while not l1 is None or not l2 is None:
            if l1 is None:
                l3.next = l2
                l2 = l2.next
                l3 = l3.next
                continue
            if l2 is None:
                l3.next = l1
                l1 = l1.next
                l3 = l3.next
                continue
            if l1.val < l2.val:
                l3.next = l1
                l1 = l1.next
                l3 = l3.next
            else:
                l3.next = l2
                l2 = l2.next
                l3 = l3.next
        return head
