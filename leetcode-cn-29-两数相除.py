#!/usr/bin/env python
# -*- coding:utf-8 -*-
from typing import *

class Solution:

    def getBinary(self, num):
        ans = ''
        if num > 0:
            ans = self.getBinary(int(num / 2))
            ans = ans + str(num % 2)
        return ans

    def getDecimalism(self, str_num):
        step = 1
        ans = 0
        for i in range(len(str_num) - 1, -1 ,-1):
            ans = ans + int(str_num[i]) * step
            step = step * 2
        return ans

    def minus(self, str_num1, str_num2):
        int_num1 = self.getDecimalism(str_num1)
        int_num2 = self.getDecimalism(str_num2)
        if int_num1 < int_num2:
            return 0, 0
        return 1, self.getBinary(int_num1 - int_num2)

    def divide(self, dividend: int, divisor: int) -> int:
        flag = 0
        if dividend < 0:
            flag = flag ^ 1
            dividend = abs(dividend)
        if divisor < 0:
            flag = flag ^ 1
            divisor = abs(divisor)
        str_dividend = self.getBinary(dividend)
        str_divisor = self.getBinary(divisor)
        ans = ''
        index = 0
        temp = ''
        while index < len(str_dividend):
            temp = temp + str_dividend[index]
            res, remainder = self.minus(temp, str_divisor)
            if res != 0:
                temp = remainder
            ans = ans + str(res)
            index = index + 1

        if flag:
            int_ans = -1 * self.getDecimalism(ans)
        else:
            int_ans = self.getDecimalism(ans)
        if int_ans > 2147483647:
            return 2147483647
        else:
            return int_ans

solution = Solution()
print(solution.divide(-1021989372, -82778243))
