package main

import (
	"fmt"
)

func setZeroes(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				for k := 0; k < len(matrix); k++ {
					if matrix[k][j] != 0 {
						matrix[k][j] = -65728324
					}
				}
				for k := 0; k < len(matrix[i]); k++ {
					if matrix[i][k] != 0 {
						matrix[i][k] = -65728324
					}
				}
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == -65728324 {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	nums := [][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}}
	setZeroes(nums)
	fmt.Println(nums)
}
