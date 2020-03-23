package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := 0
	if m == 0 {
		return 0
	}
	dp := make([][]int, m, m)
	for i := 0; i < len(dp); i++ {
		if n == 0 {
			n = len(obstacleGrid[i])
		}
		dp[i] = make([]int, n, n)
	}
	if obstacleGrid[0][0] == 0 {
		dp[0][0] = 1
	}
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if i-1 >= 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j-1 >= 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	ans := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(ans))
	ans = [][]int{}
	fmt.Println(uniquePathsWithObstacles(ans))
}
