#!/usr/bin/env python
# -*- coding:utf-8 -*-

class Solution:

    def intval(self, s) -> int:
        if s == ",":
            return 0
        return int(s)

    def check(self, nums) -> bool:
        for i in range(0, len(nums)):
            if nums[i] == 0:
                return False
        return True

    def examine(self, mark, x, y) -> bool:
        dic_sudo = [0] * 10
        for i in range(0, 10):
            dic_sudo[mark[x][i]] = 1
        if not self.check(dic_sudo)):
            return False
        for i in range(0, 10):
            dic_sudo[mark[i][y]] = 1
        if not self.check(dic_sudo)):
            return False
        x = x - x % 3
        y = y - y % 3
        dic_sudo = [0] * 10
        for i in range(x, x + 3):
            for j in range(y, y + 3):
                dic_sudo[mark[i][j]] == 1
        if not self.check(dic_sudo):
            return False
        return True

    def isValidSudoku(self, board: List[List[str]]) -> bool:
        mark = []
        for i in range(len(board)):
            mark.append(map(int,board[i]);
        return mark
        
solution = new Solution
print(solution.isValidSudoku)
