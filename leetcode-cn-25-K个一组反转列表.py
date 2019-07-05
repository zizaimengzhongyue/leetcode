#!/usr/bin/env python
# -*- coding:utf-8 -*-

class Solution:
    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        q = []
        cur = head
        while not cur is None:
            q.append(cur.val)
            cur = cur.next
        index = 1
        ln = len(q)
        ans = []
        while index * k <= ln:
            for i in range(index * k - 1, index * k - k - 1, -1):
                ans.append(q[i])
            index = index + 1
        for i in range(index * k - k, ln):
            ans.append(q[i])
        if len(ans) == 0:
            return None
        ans_head = ListNode(ans[0])
        cur = ans_head
        for i in range(1, len(ans)):
            temp = ListNode(ans[i])
            cur.next = temp
            cur = temp
        return ans_head
