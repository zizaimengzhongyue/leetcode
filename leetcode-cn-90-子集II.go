package main

import (
	"fmt"
	"sort"
	"strconv"
)

func getKey(tmp []int) string {
	str := "-"
	for _, v := range tmp {
		str += "-" + strconv.Itoa(v)
	}
	return str
}

func dfs(nums []int, tmp []int, index int, start int, ans [][]int, dict map[string]bool) [][]int {
	key := getKey(tmp[0:index])
	if !dict[key] {
		dict[key] = true
		ans = append(ans, append([]int{}, tmp[0:index]...))
	}
	for i := start; i < len(nums); i++ {
		if len(tmp) <= index {
			tmp = append(tmp, nums[i])
		} else {
			tmp[index] = nums[i]
		}
		ans = dfs(nums, tmp, index+1, i+1, ans, dict)
	}
	return ans
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	return dfs(nums, []int{}, 0, 0, [][]int{}, map[string]bool{})
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
	nums = []int{}
	fmt.Println(subsetsWithDup(nums))
	nums = []int{1, 2, 3, 4, 5}
	fmt.Println(subsetsWithDup(nums))
}
