#!/usr/bin/env python
# -*- coding:utf-8 -*-
from typing import List

class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        ln = 0
        for i in range(0, len(nums)):
            if i - 1 >= 0 and nums[i] == nums[i-1]:
                continue
            nums[ln] = nums[i]
            ln = ln + 1
        return ln

nums = [1,1,2]
solution = Solution()
print(solution.removeDuplicates(nums))
