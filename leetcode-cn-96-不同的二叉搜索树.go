package main

import (
	"fmt"
)

func numTrees(n int) int {
	dp := []int{1}
	for i := 1; i <= n; i++ {
		dp = append(dp, 0)
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}
	return dp[n]
}

func main() {
	num := 3
	fmt.Println(numTrees(num))
	num = 2
	fmt.Println(numTrees(num))
	num = 5
	fmt.Println(numTrees(num))
}
