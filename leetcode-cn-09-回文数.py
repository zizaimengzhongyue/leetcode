class Solution:
    def isPalindrome(self, x):
        t = x
        x = x if x > 0 else x * -1
        ans = 0
        while x > 0:
            ans = ans * 10 + x % 10
            x //= 10
        return True if ans == t else False 
        
x = 123
#x = -123
x = 121
solution = Solution()
print(solution.isPalindrome(x))
