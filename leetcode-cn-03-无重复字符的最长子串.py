class Solution:

    def lengthOfLongestSubstring(self, s):
        last = {}
        ans = 0
        for i in range(len(s)):
            temp = i + 1 if not last.has_key(s[i]) else i - last[s[i]] + 1
            ans = ans if ans > temp else temp
            last[s[i]] = i
        return ans
        
solution = Solution()
str_1 = "dvdf"
str_2 = "pwwkew"
str_3 = "bbbbb"
ln = solution.lengthOfLongestSubstring(str_1)
print(ln)
ln = solution.lengthOfLongestSubstring(str_2)
print(ln)
ln = solution.lengthOfLongestSubstring(str_3)
print(ln)
