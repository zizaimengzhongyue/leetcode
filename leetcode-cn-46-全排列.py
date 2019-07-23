#!/usr/bin/env python
# -*- coding:utf-8 -*-

class Solution:

    def dfs(self, nums, ans, temp):
        if len(temp) == len(nums):
            ans.append(temp.copy())
            return
        for i in range(0, len(nums)):
            if self.mark[i] == 0:
                temp.append(nums[i])
                self.mark[i] = 1
                self.dfs(nums, ans, temp)
                self.mark[i] = 0
                temp.pop()
        return

    
    def permute(self, nums: List[int]) -> List[List[int]]:
        ans = []
        self.mark = []
        for i in range(0, len(nums)):
            self.mark.append(0)
        for i in range(0, len(nums)):
            temp = []
            temp.append(nums[i])
            self.mark[i] = 1
            self.dfs(nums, ans, temp)
            self.mark[i] = 0
            temp.pop()
        return ans
