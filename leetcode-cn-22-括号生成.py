from typing import List
class Solution:

    def dfs(self, str, left, right):
        if left + right == self.n and left == right:
            self.ans.append(str)
            return
        if left + right == self.n:
            return
        if left > right:
            self.dfs(str + ')', left, right + 1)
        self.dfs(str + '(', left + 1, right)

    def generateParenthesis(self, n: int) -> List[str]:
        self.n = n * 2
        self.ans = []
        left = right = 0
        self.dfs('', left, right)
        return self.ans

n = 6
solution = Solution()
print(solution.generateParenthesis(n))
