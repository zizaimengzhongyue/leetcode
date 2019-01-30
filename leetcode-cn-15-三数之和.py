class Solution:
    def threeSum(self, nums):
        dic = {}
        nums = sorted(nums, reverse = True)
        ln = len(nums)
        for i in range(0, ln):
            if nums[i] < 0:
                break
            for j in range(i + 1, ln):
                key_i = nums[i] + nums[j]
                if key_i <= 0 and nums[j] < 0:
                    break
                if not key_i in dic:
                    dic[key_i] = {}
                key_j = str(nums[i]) + "_" + str(nums[j])
                dic[key_i][key_j] = [i, j]
        ans = []
        ans_keys = {}
        for i in range(0, len(nums)):
            key = 0 - nums[i]
            if key in dic:
                for x in dic[key]:
                    if dic[key][x][0] != i and dic[key][x][1] != i:
                        key_0 = nums[dic[key][x][0]]
                        key_1 = nums[dic[key][x][1]]
                        key_2 = nums[i]
                        if key_1 < key_2:
                            temp = key_1
                            key_1 = key_2
                            key_2 = temp
                        key_temp = str(key_0) + '_' + str(key_1) + '_' + str(key_2)
                        if not key_temp in ans_keys:
                            ans_keys[key_temp] = 1
                            ans.append([key_0, key_1, key_2])
                dic.pop(key)
        return ans
