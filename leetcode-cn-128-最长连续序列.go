package main

import "fmt"
import "strconv"

func longestConsecutive(nums []int) int {
	max := 0
	dp := map[string]int{}
	mark := map[string]int{}
	for _, v := range nums {
		strV := strconv.Itoa(v)
		dp[strV] = 1
	}
	for _, v := range nums {
		tmp := v
		count := 0
		sTmp := strconv.Itoa(tmp)
		if mark[sTmp] != 0 {
			continue
		}
		for mark[sTmp] == 0 && dp[sTmp] == 1 {
			mark[sTmp] = 1
			count = count + 1
			tmp = tmp - 1
			sTmp = strconv.Itoa(tmp)
		}
		tmp = v + 1
		sTmp = strconv.Itoa(tmp)
		for mark[sTmp] == 0 && dp[sTmp] == 1 {
			mark[sTmp] = 1
			count = count + 1
			tmp = tmp + 1
			sTmp = strconv.Itoa(tmp)
		}
		if count > max {
			max = count
		}
	}
	return max
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
	nums = []int{0, -1}
	fmt.Println(longestConsecutive(nums))
	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
	nums = []int{-2, -3, -3, 7, -3, 0, 5, 0, -8, -4, -1, 2}
	fmt.Println(longestConsecutive(nums))
}
