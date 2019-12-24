package main

import (
	"fmt"
)

func dfs(candidates []int, nums []int, begin int, cur int, target int, ans [][]int) [][]int {
}

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	temp := []int{}
	cur := 0
	dfs(candidates, temp, 0, 0, target, ans)
	return ans
}

func main() {
	nums := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(nums, target))

	nums = []int{2, 3, 5}
	target = 8
	fmt.Println(combinationSum(nums, target))
}
