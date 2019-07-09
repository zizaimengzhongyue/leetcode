#!/usr/bin/env python
# -*- coding:utf-8 -*-
class Solution:

    def setNext(self, needle):
        ln = len(needle)
        j = 0
        k = -1
        self.next = []
        self.next.append(k)
        while j < ln:
            if k == -1 or needle[j] == needle[k]:
                k = k + 1
                j = j + 1
                self.next.append(k)
            else:
                k = self.next[k]

    def kmp(self, haystack, needle):
        self.setNext(needle)
        i = 0
        j = 0
        ln_h = len(haystack)
        ln_n = len(needle)
        while i < ln_h and j < ln_n:
            if j == -1 or haystack[i] == needle[j]:
                i = i + 1
                j = j + 1
            else:
                j = self.next[j]
        if j == ln_n:
            return i - j
        else :
            return -1

    def strStr(self, haystack: str, needle: str) -> int:
        if len(needle) == 0:
            return 0
        return self.kmp(haystack, needle)

haystack = 'BBC ABCDAB ABCDABCDABDE'
needle = 'ABCDABD'
solution = Solution()
print(solution.strStr(haystack, needle))
