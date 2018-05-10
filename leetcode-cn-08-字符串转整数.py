class Solution:
    def myAtoi(self, str):
        ans = 0
        flag = 0
        symbol = 1
        for i in range(len(str)):
            if flag == 0 and str[i] == ' ':
                continue
            if flag == 0 and str[i] == '-':
                flag = -1
                symbol = -1
            elif flag == 0 and str[i] == '+':
                flag = -1
                symbol = 1
            elif str[i] < '0' or str[i] > '9':
                flag = 1
            elif str[i] >= '0' and str[i] <= '9':
                flag = -1
            if flag > 0:
                break
            if str[i] >= '0' and str[i] <= '9':
                ans = ans * 10 + (ord(str[i]) - ord('0'))
        ans *= symbol
        ans = -2147483648 if ans < -2147483648 else ans
        ans = 2147483647 if ans > 2147483647 else ans
        return ans

s = "42"
s = "   -42"
s = "4193 with words"
s = "words and 987"
#s = "-91283472332"
#s = "--42"
#s = "   +0 123"
solution = Solution()
print(solution.myAtoi(s))
