#!/usr/bin/env python
# -*- coding:utf-8 -*-
from typing import *

class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        ln = len(nums)
        flag = False
        for i in range(ln - 1, -1, -1):
            for j in range(i, ln):
                if nums[j] > nums[i]:
                    temp = nums[j]
                    nums[j] = nums[i]
                    nums[i] = temp
                    flag = True
                    break
            if flag:
                break
            for j in range(i, ln - 1):
                if nums[j] > nums[j + 1]:
                    temp = nums[j]
                    nums[j] = nums[j + 1]
                    nums[j + 1] = temp
                else:
                    break

nums = [1,2,3]
nums = [2,3,1]
nums = [3,2,1]
nums = [4,2,0,2,3,2,0]
#nums = [1,3,2]
#nums = [3,2,1]
solution = Solution()
solution.nextPermutation(nums)
print(nums)
