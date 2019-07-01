class Solution:

    def fourSum(self, nums, target):
        nums.sort()
        ans = []
        ans_dic = {}
        left_dic = {}
        right_dic = {}
        nums_len = len(nums)
        for i in range (0, nums_len):
            for j in range(i + 1, nums_len):
                str_key = str(i) + '_' + str(j)
                left_dic[str_key] = nums[i] + nums[j]
        for i in range(0, nums_len):
            for j in range(i + 1, nums_len):
                str_key = str(i) + '_' + str(j)
                int_value = str(nums[i] + nums[j])
                if int_value in right_dic:
                    right_dic[int_value].append(str_key)
                else:
                    right_dic[int_value] = [str_key]
        left_len = len(left_dic)
        right_len = len(right_dic)
        for left_key in left_dic:
            cur_target = target - left_dic[left_key]
            cur_target_key = str(cur_target)
            if cur_target_key in right_dic:
                for i in range(0, len(right_dic[cur_target_key])):
                    right_key = right_dic[cur_target_key][i]
                    ans_key = left_key + "_" + right_key
                    list_ans_key = ans_key.split("_")
                    x0 = int(list_ans_key[0])
                    x1 = int(list_ans_key[1])
                    x2 = int(list_ans_key[2])
                    x3 = int(list_ans_key[3])
                    if x0 >= x1 or x1 >= x2 or x2 >= x3:
                        continue
                    ans_key = str(nums[x0]) + '_' + str(nums[x1]) + '_' + str(nums[x2]) + '_' + str(nums[x3])
                    if ans_key in ans_dic:
                        continue
                    ans_dic[ans_key] = 1
                    ans.append([nums[x0], nums[x1], nums[x2], nums[x3]])
                        
        return ans
        
nums = [1, 0, -1, 0, -2, 2]
nums = [0, 0, 0, 0]
nums = [-3,-2,-1,0,0,1,2,3]
target = 0
solution = Solution()
ans = solution.fourSum(nums, target)
print(ans)
