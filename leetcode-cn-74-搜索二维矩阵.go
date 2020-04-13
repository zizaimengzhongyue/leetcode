package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	if len(matrix) == 1 && len(matrix[0]) == 1 {
		return matrix[0][0] == target
	}
	row := 0
	left := 0
	right := len(matrix) - 1
	flag := 0
	for left < right {
		mid := (left + right) / 2
		if left == right-1 && flag == 0 {
			flag++
		} else if left == right-1 && flag == 1 {
			mid++
			flag++
		} else if left == right-1 {
			return false
		}
		if (matrix[mid][0] <= target && mid == len(matrix)-1) || (matrix[mid][0] <= target && matrix[mid+1][0] > target) {
			row = mid
			break
		}
		if matrix[mid][0] > target {
			right = mid
		}
		if matrix[mid][0] <= target {
			left = mid
		}
	}
	left = 0
	right = len(matrix[0]) - 1
	if left == right && matrix[row][left] == target {
		return true
	}
	flag = 0
	for left < right {
		mid := (left + right) / 2
		if left == right-1 && flag == 0 {
			flag++
		} else if left == right-1 && flag == 1 {
			mid++
			flag++
		} else if left == right-1 {
			return false
		}
		if matrix[row][mid] == target {
			return true
		}
		if matrix[row][mid] > target {
			right = mid
		}
		if matrix[row][mid] <= target {
			left = mid
		}
	}
	return false
}

func main() {
	nums := [][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 50}}
	target := 3
	fmt.Println(searchMatrix(nums, target))
	nums = [][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 50}}
	target = 50
	fmt.Println(searchMatrix(nums, target))
	nums = [][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 50}}
	target = 51
	fmt.Println(searchMatrix(nums, target))
	nums = [][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 50}}
	target = -1
	fmt.Println(searchMatrix(nums, target))
	nums = [][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 50}}
	target = 13
	fmt.Println(searchMatrix(nums, target))
	nums = [][]int{[]int{1, 3}}
	target = 3
	fmt.Println(searchMatrix(nums, target))
	nums = [][]int{[]int{1}, []int{3}}
	target = 3
	fmt.Println(searchMatrix(nums, target))
}
