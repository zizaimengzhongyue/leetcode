#!/usr/bin/env python
# -*- coding:utf-8 -*-
from typing improt List

class Solution:
    def swapPairs(self, head: ListNode) -> ListNode:
        if head is None or head.next is None:
            return head
        first = head
        second = head.next
        after = second.next
        first.next = after
        second.next = first
        head = second
        
        before = first
        cur = after
        while not cur is None:
            first = cur
            second = cur.next
            if second is None:
                break
            after = None
            if not second is None:
                after = second.next
            before.next = second
            first.next = after
            second.next = first

            before = first
            cur = after
        
        return head

