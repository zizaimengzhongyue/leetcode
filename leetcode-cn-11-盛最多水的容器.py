class Solution:
    def maxArea(self, height):
        sort_height = []
        for i in range(0, len(height)):
            sort_height.append((i,height[i]))
        sorted_height = sorted(sort_height, key=lambda s:s[1], reverse=True)
        max_area = 0
        left = sorted_height[0][0]
        right = sorted_height[0][0]
        for i in range(1, len(sorted_height)):
            temp_height = sorted_height[i][1]
            temp_index = sorted_height[i][0]
            temp_width = abs(temp_index - left) if abs(temp_index - left) > abs (temp_index - right) else abs(temp_index - right)
            temp_area = temp_width * temp_height
            max_area = temp_area if temp_area > max_area else max_area
            left = temp_index if temp_index < left else left
            right = temp_index if temp_index > right else right
        return max_area

height = [5,7,6,4,3,2,9]
height = [0,2]
#height = [1,1]
solution = Solution()
print(solution.maxArea(height))
