package main

import (
	"fmt"
)

func dfs(candidates []int, nums []int, ln int, begin int, cur int, target int, ans [][]int) [][]int {
	if cur == target {
		return append(ans, append([]int{}, nums[0:ln]...))
	}
	for i := begin; i < len(candidates); i++ {
		if cur+candidates[i] > target {
			continue
		}
		if len(nums) <= ln {
			nums = append(nums, candidates[i])
		} else {
			nums[ln] = candidates[i]
		}
		ans = dfs(candidates, nums, ln+1, i, cur+candidates[i], target, ans)
	}
	return ans
}

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	temp := []int{}
	ans = append(ans, dfs(candidates, temp, 0, 0, 0, target, ans)...)
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
