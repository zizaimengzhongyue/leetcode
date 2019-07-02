class Solution:
    def isValid(self, s):
        queue = []
        index = 0
        for i in range(0, len(s)):
            if s[i] == '[' or s[i] == '{' or s[i] == '(':
                queue.append(s[i])
                index = index + 1
            elif index == 0:
                return False
            if s[i] == ']':
                if queue[index - 1] == '[':
                    index = index - 1
                    queue.pop()
                else:
                    return False
            elif s[i] == '}':
                if queue[index - 1] == '{':
                    index = index - 1
                    queue.pop()
                else:
                    return False
            elif s[i] == ')':
                if queue[index - 1] == '(':
                    index = index - 1
                    queue.pop()
                else:
                    return False
        if len(queue) == 0:
            return True
        else:
            return False

str = '('
solution = Solution()
print(solution.isValid(str))
