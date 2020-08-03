package main

import (
	"fmt"
)

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	dp := [][]int{}
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		tmp := []int{}
		for j := 0; j < n; j++ {
			tmp = append(tmp, 0)
		}
		dp = append(dp, tmp)
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
		}
	}
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = dp[i][j-1] + 1
			}
		}
	}
	max := 0
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
		}
	}
	return max
}

func main() {
	bytes := [][]byte{[]byte{'1', '0', '1', '0', '0'}, []byte{'1', '0', '1', '1', '1'}, []byte{'1', '1', '1', '1', '1'}, []byte{'1', '0', '0', '1', '0'}}
	fmt.Println(maximalRectangle(bytes))
}
