package main

import "fmt"

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	n := 0
	dp := make([][]int, len(grid), len(grid))
	for i := 0; i < len(grid); i++ {
		if n == 0 {
			n = len(grid[i])
		}
		for j := 0; j < len(grid[i]); j++ {
			dp[i] = make([]int, n, n)
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
				continue
			}
			tmp1 := 0
			tmp2 := 0
			if i-1 >= 0 {
				tmp1 = grid[i][j] + dp[i-1][j]
			}
			if j-1 >= 0 {
				tmp2 = grid[i][j] + dp[i][j-1]
			}
			if tmp1 == 0 {
				dp[i][j] = tmp2
				continue
			}
			if tmp2 == 0 {
				dp[i][j] = tmp1
				continue
			}
			if tmp1 < tmp2 {
				dp[i][j] = tmp1
			} else {
				dp[i][j] = tmp2
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid))
}
