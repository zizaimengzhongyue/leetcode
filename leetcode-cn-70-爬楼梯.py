#!/usr/bin/env python
# -*- coding:utf-8 -*-

from typing import *

class Solution:
    
    def climbStairs(self, n: int) -> int:
        ans = [1, 1]
        for i in range(2, n + 1):
            ans.append(ans[i - 1] + ans[i - 2])
        return ans[n]

solution = Solution()
print(solution.climbStairs(3))
