#!/usr/bin/env python
# -*- coding:utf-8 -*-

class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        temp = 0
        for i in range(len(digits) - 1, -1, -1):
            if i == len(digits) - 1:
                digits[i] = digits[i] + 1
            else:
                digits[i] = digits[i] + temp
            if digits[i] >= 10:
                digits[i] = digits[i] % 10
                temp = 1
            else:
                temp = 0

        ans = []
        if temp > 0:
            ans = [temp] + digits
        else:
            ans = digits
        return ans
            
