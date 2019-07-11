#!/usr/bin/env python
# -*- coding:utf-8 -*-

class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        ln = len(nums)
        temp = nums[ln - 1]
        for i in range(ln - 1, -1, -1):
            if nums[i] < temp:
                nums[ln - 1] = nums[i]
                nums[i] = temp
                break

