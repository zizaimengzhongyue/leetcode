#!/usr/bin/env python
# -*- coding:utf-8 -*-

class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        ln = 0
        bak = ln
        for i in range(0, len(s)):
            if s[i] == ' ':
                if ln != 0:
                    bak = ln
                ln = 0
            else:
                ln = ln + 1
        if ln == 0:
            ln = bak
        return ln
