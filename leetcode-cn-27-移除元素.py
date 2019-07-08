#!/usr/bin/env python
# -*- coding:utf-8 -*-
class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        ln = 0
        for i in range(0, len(nums)):
            if nums[i] == val:
                continue
            nums[ln] = nums[i]
            ln = ln + 1
        return ln

