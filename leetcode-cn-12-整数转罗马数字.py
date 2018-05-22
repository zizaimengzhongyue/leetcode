class Solution:
    def intToRoman(self, num):
        mod = [[1000,'M'], [500,'D'], [100,'C'], [50,'L'], [10,'X'], [5,'V'], [1,'I']]
        str = ''

        k = num // 1000
        if k:
            str = str + k * 'M'
        num = num % 1000

        k = num // 100
        if k == 9:
            str = str + 'CM'
            num = num % 100
        elif k == 4:
            str = str + 'CD'
            num = num % 100

        k = num // 500
        if k:
            str = str + k * 'D'
            num = num % 500

        k = num // 100
        if k:
            str = str + k * 'C'
            num = num % 100

        k = num // 10
        if k == 9:
            str = str + 'XC'
            num = num % 10
        elif k == 4:
            str = str + 'XL'
            num = num % 10

        k = num // 50
        if k:
            str = str + k * 'L'
            num = num % 50

        k = num // 10
        if k:
            str = str + k * 'X'
            num = num % 10

        if num == 9:
            str = str + 'IX'
            num = num - 9
        elif num == 4:
            str = str + 'IV'
            num = num - 4

        if num >= 5:
            str = str + 'V'
            num = num - 5

        str = str + num * 'I'

        return str

num = 1994
num = 58
num = 9
num = 4
num = 3
num = 100
num = 5
solution = Solution()
print(solution.intToRoman(num))
