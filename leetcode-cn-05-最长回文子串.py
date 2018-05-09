class Solution:
    def longestPalindrome(self, s):
        if len(s) == 0:
            return 0
        ans = 1
        ans_l = 0
        ans_r = 1
        for i in range(len(s)):
            temp = 1
            l = i - 1
            r = i + 1
            while l >= 0 and r < len(s) and s[l] == s[r]:
                temp += 2
                if temp > ans:
                    ans = temp
                    ans_l = l
                    ans_r = r + 1
                l -= 1
                r += 1
            temp = 0
            l = i
            r = i + 1
            while l >= 0 and r < len(s) and s[l] == s[r]:
                temp += 2
                if temp > ans:
                    ans = temp
                    ans_l = l
                    ans_r = r + 1
                l -= 1
                r += 1
        return s[ans_l:ans_r]  


str = "babad"
#str = "cbbd"
solution = Solution()
print(solution.longestPalindrome(str))
