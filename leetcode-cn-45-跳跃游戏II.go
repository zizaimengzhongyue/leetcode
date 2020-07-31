package main

import (
	"fmt"
)

func jump(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums) && j <= i+nums[i]; j++ {
			if dp[i]+1 < dp[j] || dp[j] == 0 {
				dp[j] = dp[i] + 1
			}
		}
	}
	return dp[len(nums)-1]
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}
