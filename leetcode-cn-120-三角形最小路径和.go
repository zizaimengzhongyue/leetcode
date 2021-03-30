package main

import (
	"fmt"
)

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if triangle[i+1][j] < triangle[i+1][j+1] {
				triangle[i][j] += triangle[i+1][j]
			} else {
				triangle[i][j] += triangle[i+1][j+1]
			}
		}
	}
	return triangle[0][0]
}

func main() {
	arr := [][]int{}
	arr = append(arr, []int{2})
	arr = append(arr, []int{3, 4})
	arr = append(arr, []int{6, 5, 7})
	arr = append(arr, []int{4, 1, 8, 3})
	fmt.Println(minimumTotal(arr))
}
