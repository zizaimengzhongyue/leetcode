class Solution:

    """
    返回值>0:符合要求的数字数量大于target_num
    <0:小于target_num
    =0:等于target_num
    """
    def countSmallNum(self, num, target, target_num):
        if target_num < 0:
            return 1
        if len(num) == 0 and target_num == 0:
            return 0
        elif len(num) == 0:
            return -1
        if target_num > len(num):
            return -1
        if target_num == 0 and num[0] < target:
            return 1
        if target_num == 0 and num[0] >= target:
            return 0
        if target_num == len(num) and num[len(num) - 1] > target:
            return -1
        if target_num == len(num) and num[len(num) - 1] <= target:
            return 0
        if num[target_num - 1] <= target and num[target_num] >= target:
            return 0
        if num[target_num - 1] > target:
            return -1
        if num[target_num] < target:
            return 1

    def findMedian(self, nums1, nums2, target):
        left = 0
        right = len(nums1) - 1
        while left <= right:
            mid = (left + right) // 2
            res = self.countSmallNum(nums2, nums1[mid], target - mid)
            if res == 0:
                return nums1[mid]
            elif res < 0:
                left = mid + 1
            elif res > 0:
                right = mid - 1
        left = 0
        right = len(nums2) - 1
        while left <= right:
            mid = (left + right) // 2
            res = self.countSmallNum(nums1, nums2[mid], target - mid)
            if res == 0:
                return nums2[mid]
            elif res < 0:
                left = mid + 1
            elif res > 0:
                right = mid - 1
        return -1

    def findMedianSortedArrays(self, nums1, nums2):
        sum = len(nums1) + len(nums2)
        if sum % 2:
            return self.findMedian(nums1, nums2, (sum + 1) // 2 - 1)
        else:
            target_1 = self.findMedian(nums1, nums2, sum // 2 - 1);
            target_2 = self.findMedian(nums1, nums2, sum // 2);
            return (target_1 + target_2) / 2


nums1 = [1]
nums2 = [3,4]
#nums1 = [1,3]
#nums2 = [2]
#nums1 = [1]
#nums2 = [2]
#nums1 = [1,2,3,5,8,9]
#nums2 = [4,6,7,10]
#nums1 = []
#nums2 = [1]
#nums1 = [1]
#nums2 = [1]
#nums1 = []
#nums2 = [1,2,3,4,5,6]
solution = Solution()
print(solution.findMedianSortedArrays(nums1,nums2))
