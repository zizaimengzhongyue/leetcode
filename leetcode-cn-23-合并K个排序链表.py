from typing import List

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def mergeKLists(self, lists: List[ListNode]) -> ListNode:
        heap_q = []
        for item in lists:
            while not item is None:
                heap_q.append(item.val)
                item = item.next
        heap_q.sort()
        if len(heap_q) == 0:
            return None
        head = ListNode(heap_q[0])
        cur = head
        for i in range(1, len(heap_q)):
            temp = ListNode(heap_q[i])
            cur.next = temp
            cur = temp
        return head

