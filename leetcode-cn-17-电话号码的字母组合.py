class Solution:
    def letterCombinations(self, digits):
	ln = len(digits)
	if ln == 0:
	    return []
	ans = []
	self.solute(ans, digits, '',  0, ln)
	return ans

    def solute(self, ans, digits, str_ans, index, ln):
	dic = {'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	key = digits[index]
	for i in range(0, len(dic[key])):
	    if index == ln - 1:
		ans.append(str_ans + dic[key][i])
	    else:
	    	self.solute(ans, digits, str_ans + dic[key][i], index + 1, ln)
		    
solution = Solution()
print(solution.letterCombinations("23"))
