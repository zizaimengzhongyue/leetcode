class Solution:
    def convert(self, s, n):
        if n <= 1:
            return s
        ln = len(s)
        str = ""
        m = ln // (2 * n - 2) * (n - 1)
        mod = ln - ln // (2 * n - 2) * (2 * n - 2)
        if mod == 0:
            m = m
        elif mod < n:
            m += 1
        else:
            m += mod - n + 1
        for i in range(0, n):
            for j in range(0, m):
                index = -1
                if j % (n - 1) == 0:
                    index = j // (n - 1) * (2 * n - 2) + i
                elif j % (n - 1) == n - i - 1:
                    index = j // (n - 1) * (2 * n - 2) + j % (n - 1) + n - 1
                if index >= 0 and index < ln:
                    str += s[index]
        return str

s = "PAYPALISHIRING"
numRows = 3
#s = "PAYPALISHIRING"
#numRows = 4
s = "ikicnhlvnsnklobqkfoifwsdzfgoueicotgearzqloawcbomgcwnyosztzoqdyutrbxsuzignicqkresvbgwravnzjdekzabaxvvwddmoheaaiuwlcuddpqebvjlwjmxiiuhnztdgzfhhaftczyhoqmrnxcjirisbezqduuktztdbywezwlnmzfzwjkgatzdxaubvolqpgtrerxdpksbmzckjnrlgqpdqjgnztgqzbjftgkktdmfeqppdgdlysrrdxgermuqogyjmithlhmsychdkkpleicjfinyxkrlqpuexpmlxydvcmapxradzcnfmafjmwwcnqhlityftdcypenownsvrfcdhipuglwuefhmahtwkxuzbhqxuivnnxquhhootnkxstlvaermblnesdcolxbawtvebluuagttbeqbihnlucpmgtcrsdupypvzcrvdxoyysobcxqprshfyahsvmmmkdhmhahkxykxsaavxqkokrvlucfpqtbbwkpfkyqjwcipdimaayrmsmxmredhxgkvqzbcytluvvlydxbfyurqirezvhjtptohtrjtkwngrbgwyjbwdsomjsnfbwaqqqxligeseravujrvsutynimyazdcwfsmwszeadyfwbarmumgofednqcqshuonoclgbrhepmasorgkfnblteamisesmshbwshxjqdaeftvzzfeogrinsuvkapzswaqojpxcuagzvbywhflpwveqpcdprcjoebepqcrkeyjzbkvkrganbggsiljapsqtvuilycxcajbdtxokfqzhwfbfslhhfxabtlmspkuptyuoxiacyzjxhlezylhdkjefwtxlfucyuxorhotipffysjyqwhtisynrtley"
numRows = 158
solution = Solution()
print(solution.convert(s, numRows))
