#!/usr/bin/env python
# -*- coding:utf-8 -*-

class Solution:

    def addBinary(self, a: str, b: str) -> str:
        ln_a = len(a)
        ln_b = len(b)
        index_a = ln_a - 1
        index_b = ln_b - 1
        ans = ''
        temp = 0
        while index_a >= 0 or index_b >= 0:
            if index_a >= 0:
                int_a = int(a[index_a])
            else:
                int_a = 0
            if index_b >= 0:
                int_b = int(b[index_b])
            else:
                int_b = 0
            int_ans = int_a + int_b + temp
            if int_ans >= 2:
                int_ans = int_ans % 2
                temp = 1
            else:
                temp = 0
            ans =  ans + str(int_ans)
            index_a = index_a - 1
            index_b = index_b - 1
        if temp > 0:
            ans = ans + str(temp)
        real_ans = ''
        for i in range(len(ans) - 1, -1, -1):
            real_ans = real_ans + ans[i]
        return real_ans
