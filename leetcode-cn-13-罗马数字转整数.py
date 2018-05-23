class Solution:
    def romanToInt(self, s):
        ans = 0
        ln = len(s)
        i = 0
        while i < ln:
            if s[i] == 'M':
                ans += 1000
                i += 1
                continue
            if s[i] == 'D':
                ans += 500
                i += 1
                continue
            if s[i] == 'C':
                if i + 1 < ln:
                    if s[i + 1] == 'M':
                        ans += 900
                        i += 1
                    elif s[i + 1] == 'D':
                        ans += 400
                        i += 1
                    else:
                        ans += 100
                else:
                    ans += 100
                i += 1
                continue
            if s[i] == 'L':
                ans += 50 
                i += 1
                continue
            if s[i] == 'X':
                if i + 1 < ln:
                    if s[i + 1] == 'C':
                        ans += 90
                        i += 1
                    elif s[i + 1] == 'L':
                        ans += 40
                        i += 1
                    else:
                        ans += 10
                else:
                    ans += 10
                i += 1
                continue
            if s[i] == 'V':
               ans += 5
               i += 1
               continue
            if s[i] == 'I':
                if i + 1 < ln:
                    if s[i + 1] == 'X':
                        ans += 9
                        i += 1
                    elif s[i + 1] == 'V':
                        ans += 4
                        i += 1
                    else:
                        ans += 1
                else:
                    ans += 1
                i += 1
                continue

        return ans

s = "MCMXCIV"
s = "LVIII"
s = "IX"
s = "V"
s = "III"
solution = Solution()
print(solution.romanToInt(s))
