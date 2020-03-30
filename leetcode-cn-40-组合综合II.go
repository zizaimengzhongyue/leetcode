package main

import (
	"fmt"
	"sort"
	"strconv"
)

func dfs(candidates []int, nums []int, ln int, begin int, cur int, target int, ans [][]int, mark map[string]bool) [][]int {
	if cur == target {
		key := ""
		for _, v := range nums[0:ln] {
			key = key + "-" + strconv.Itoa(v)
		}
		if mark[key] {
			return ans
		} else {
			mark[key] = true
		}
		return append(ans, append([]int{}, nums[0:ln]...))
	}
	for i := begin; i < len(candidates); i++ {
		if cur+candidates[i] > target {
			return ans
		}
		if len(nums) <= ln {
			nums = append(nums, candidates[i])
		} else {
			nums[ln] = candidates[i]
		}
		ans = dfs(candidates, nums, ln+1, i+1, cur+candidates[i], target, ans, mark)
	}
	return ans
}

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	ans := [][]int{}
	temp := []int{}
	sort.Ints(candidates)
	mark := map[string]bool{}
	ans = dfs(candidates, temp, 0, 0, 0, target, ans, mark)
	return ans
}

func main() {
	nums := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum2(nums, target))

	nums = []int{2, 3, 5}
	target = 8
	fmt.Println(combinationSum2(nums, target))

	nums = []int{10, 1, 2, 7, 6, 1, 5}
	target = 8
	fmt.Println(combinationSum2(nums, target))
}
