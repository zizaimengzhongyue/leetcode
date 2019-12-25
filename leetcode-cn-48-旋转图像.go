package main

import (
	"fmt"
)

func rotate(nums [][]int) [][]int {
	n := len(nums)
	for i := 0; i <= n/2; i++ {
		for j := i; j <= n-i-2; j++ {
			x := i
			y := j
			nextX := y
			nextY := n - x - 1
			temp := nums[x][y]
			nextTemp := nums[nextX][nextY]
			for nextX != i || nextY != j {
				nextTemp = nums[nextX][nextY]
				nums[nextX][nextY] = temp
				temp = nextTemp
				x = nextX
				y = nextY
				nextX = y
				nextY = n - x - 1
			}
			nums[nextX][nextY] = temp
		}
	}
	return nums
}

func main() {
	nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(rotate(nums))
	nums = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	fmt.Println(rotate(nums))
}
