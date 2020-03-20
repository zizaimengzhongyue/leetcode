package main

import "fmt"

func uniquePaths(m int, n int) int {
	dp := make([][]int, m+1, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1, n+1)
	}
	dp[0][1] = 1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m][n]
}

func main() {
	m := 3
	n := 2
	fmt.Println(uniquePaths(m, n))
	m = 1
	n = 1
	fmt.Println(uniquePaths(m, n))
	m = 7
	n = 3
	fmt.Println(uniquePaths(m, n))
}
