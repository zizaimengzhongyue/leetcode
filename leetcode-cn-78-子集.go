package main

import "fmt"

func dfs(tmp []int, index int, nums []int, count int) [][]int {
	ans := [][]int{}
	if index == len(nums) {
		return ans
	}
	ans = append(ans, []int{})
	for i := 0; i < count; i++ {
		ans[len(ans)-1] = append(ans[len(ans)-1], tmp[i])
	}
	for i := index + 1; i < len(nums); i++ {
		tmp[count] = nums[i]
		ans = append(ans, dfs(tmp, i, nums, count+1)...)
	}
	return ans
}

func subsets(nums []int) [][]int {
	ans := [][]int{[]int{}}
	tmp := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		tmp[0] = nums[i]
		ans = append(ans, dfs(tmp, i, nums, 1)...)
	}
	return ans
}

func main() {
	nums := []int{3, 1, 2}
	fmt.Println(subsets(nums))
	nums = []int{}
	fmt.Println(subsets(nums))
	nums = []int{3}
	fmt.Println(subsets(nums))
	nums = []int{3, 1, 2, 4, 5}
	fmt.Println(subsets(nums))
}
