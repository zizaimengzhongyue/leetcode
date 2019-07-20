from typing import *

class Solution:

    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        index = 0
        index_1 = 0
        index_2 = 0
        while index_1 < m + n or index_2 < n:
            if index_1 >= m + index_2:
                nums1.insert(index, nums2[index_2])
                index_2 = index_2 + 1
                index_1 = index_1 + 1
            elif index_2 >= n:
                index_1 = index_1 + 1
            elif nums2[index_2] < nums1[index_1]:
                nums1.insert(index, nums2[index_2])
                index_2 = index_2 + 1
                index_1 = index_1 + 1
            else:
                index_1 = index_1 + 1
            index = index + 1
        i = m + n
        ln = len(nums1)
        while i < ln:
            nums1.pop(m + n)
            i = i + 1

nums1 = [4,5,6,0,0,0]
nums2 = [1,2,3]
solution = Solution()
solution.merge(nums1, 3, nums2, 3)
print(nums1)
