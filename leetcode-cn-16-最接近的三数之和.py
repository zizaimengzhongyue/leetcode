class Solution:
    def threeSumClosest(self, nums, target):
        dic = {}
        ln = len(nums)
        nums = sorted(nums, reverse = True)
        gap = 999999999
        for i in range(0, ln):
            for j in range(i + 1, ln):
                for k in range(j + 1, ln):
                    temp = nums[i] + nums[j] + nums[k]
                    if abs(target - temp) < gap:
                        gap = abs(target - temp)
                        ans = temp
        return ans
