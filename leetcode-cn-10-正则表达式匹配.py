class Solution:

    def isMatch(self, s, p):
        ln_s = len(s)
        ln_p = len(p)
        dp = []
        for i in range(0, ln_p+1):
            dp.append([])
            for j in range(0, ln_s+1):
                dp[i].append(False)
        dp[ln_p][ln_s] = True
        for i in range(ln_p-1, -1, -1):
            if p[i] == '*' and dp[i+1][ln_s] == True:
                dp[i][ln_s] = True
            elif i+1<ln_p and p[i+1]=='*' and dp[i+1][ln_s] == True:
                dp[i][ln_s] = True
        for i in range(ln_p-1, -1, -1):
            if dp[i+1][ln_s] == True and ln_s-1>=0 and p[i] == s[ln_s-1]:
                dp[i][ln_s-1] = True
            elif dp[i+1][ln_s-1] == True and p[i] == '*':
                dp[i][ln_s-1] = True
            elif dp[i+1][ln_s] == True and p[i] == '.' and ln_s-1>=0:
                dp[i][ln_s-1] = True
            elif dp[i+1][ln_s] == True and p[i] == '*':
                if i-1>=0 and ln_s-1>=0 and p[i-1]==s[ln_s-1]:
                    dp[i][ln_s-1] = True
                elif i-1>=0 and p[i-1]=='.':
                    dp[i][ln_s-1]= True
        for i in range(ln_s-1, -1, -1):
            for j in range(ln_p-1, -1, -1):
                if j+1<ln_p and dp[j+1][i] == True and p[j+1] == '*':
                    dp[j][i] = True
                elif dp[j+1][i] == True and p[j] == '*':
                    dp[j][i] = True
                elif dp[j+1][i+1] == True and s[i] == p[j]:
                    dp[j][i] = True
                elif dp[j+1][i+1] == True and p[j] == '.':
                    dp[j][i] = True
                elif dp[j+1][i+1] == True and p[j] == '*':
                    if j-1>=0 and p[j-1] == s[i]:
                        dp[j][i]=True
                    elif j-1>=0 and p[j-1] == '.':
                        dp[j][i]=True
                elif dp[j][i+1] == True and p[j] == '*':
                    if j-1>=0 and p[j-1] == s[i]:
                        dp[j][i] = True
                    elif j-1>=0 and p[j-1] == '.':
                        dp[j][i] = True
        return dp[0][0]


s = "mississippi"
p = "mis*is*ip."
#s = "aasdfasdfasdfasdfas"
#p = "aasdf.*asdf.*asdf.*asdf.*s"
#s = "a"
#p = "c*a"
#s = "cbbcbaaaaaacabccbc"
#p = "a.*b*.*.*a*a*.*ba*c"
#s = "ss"
#p = "s*"
#s = "aab"
#p = "c*a*b"
#s = "ab"
#p = ".*c"
#s = "aaa"
#p = ".*"
s = ""
p = "."

solution = Solution()
print(solution.isMatch(s, p))
