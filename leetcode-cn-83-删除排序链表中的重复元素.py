#!/usr/bin/env python
# -*- coding:utf-8 -*-
class Solution:
    def deleteDuplicates(self, head):
        if head is None:
            return head
        before = head
        cur = head.next
        while not cur is None:
            if cur.val == before.val:
                cur = cur.next
            else:
                before.next = cur
                before = cur
                cur = cur.next
        before.next = cur
        return head
