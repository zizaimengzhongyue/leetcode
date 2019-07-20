#!/usr/bin/env python
# -*- coding:utf-8 -*-
class Solution:

    def transform(self, p, ans):
        if p is None:
            ans.append(p)
            return ans
        ans.append(p.val)
        self.transform(p.left, ans)
        self.transform(p.right, ans)
        return ans

    def isSameTree(self, p, q):
        ans_p = []
        ans_p = self.transform(p, ans_p)
        ans_q = []
        ans_q = self.transform(q, ans_q)
        if len(ans_p) != len(ans_q):
            return False
        ans = True
        for i in range(0, len(ans_p)):
            if ans_p[i] != ans_q[i]:
                ans = False
        return ans

