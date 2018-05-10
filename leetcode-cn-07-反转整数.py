class Solution:
    def reverse(self, x):
        res = 0
        flag = 1
        if x < 0:
            flag = -1
            x *= -1
        while x:
            res = res * 10 + x % 10
            x //= 10
        res *= flag
        return res if res >=-2147483648 and res <= 2147483647 else 0

n = 120
#n = -123
n = 12312312312312312
solution = Solution()
print(solution.reverse(n))
