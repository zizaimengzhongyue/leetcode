class Solution:
    def longestCommonPrefix(self, strs):
        if len(strs) == 0:
            return ''
        temp = strs[0]
        for i in range(1, len(strs)):
            ln_temp = len(temp)
            ln_str = len(strs[i])
            ln = ln_temp if ln_temp < ln_str else ln_str
            for j in range(0, ln):
                if temp[j] != strs[i][j]:
                    temp = temp[0:j]
                    break
            if len(temp) > ln:
                temp = temp[0:ln]
        return temp

strs = ["flower","flow","flight"]
#strs = ["dog","racecar","car"]
solution = Solution()
print(solution.longestCommonPrefix(strs))
