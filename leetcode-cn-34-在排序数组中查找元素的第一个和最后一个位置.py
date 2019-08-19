class Solution:
    def searchRange(self, nums: List[int], target:int) -> List[int]:
        begin = end = -1
        for i in range(0, len(nums)):
            if nums[i] == target:
                if begin == -1:
                    begin = i
                end = i
        return [begin, end]
