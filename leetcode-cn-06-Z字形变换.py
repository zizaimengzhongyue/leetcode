class Solution:
    def convert(self, s, n):
        str = ""
        m = len(s) // (2 * n - 2) * (n - 1)
        mod = len(s) - len(s) // (2 * n - 2) * (2 * n - 2)
        if mod == 0:
            m = m
        elif mod < n:
            m += 1
        else:
            m += mod - n
        index = 0
        for i in range(0, n - 1):
            for j in range(0, m - 1):
                if j % (n - 1) == 0 or j % (n - 1) == n - i - 1:
                    str += s[index]
                    index += 1
                else:
                    str += " "
            str += "\n"
        return str

s = "PAYPALISHIRING"
numRows = 3
s = "PAYPALISHIRING"
numRows = 4
solution = Solution()
print(solution.convert(s, numRows))
