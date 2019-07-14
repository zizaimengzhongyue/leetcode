#!/usr/bin/env python
# -*- coding:utf-8 -*-

class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        dp = []
        dp.append(nums[0])
        for i in range(1 ,len(nums)):
            if dp[i-1] + nums[i] > nums[i]:
                dp.append(dp[i-1] + nums[i])
            else:
                dp.append(nums[i])
        ans = dp[0]
        for i in range(0, len(dp)):
            if dp[i] > ans:
                ans = dp[i]
        return ans
