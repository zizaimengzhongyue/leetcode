package main

import (
	"fmt"
)

func search(nums []int, target int) bool {
	for _, v := range nums {
		if v == target {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))
	nums = []int{2, 5, 6, 0, 0, 1, 2}
	target = 3
	fmt.Println(search(nums, target))
	nums = []int{2}
	target = 3
	fmt.Println(search(nums, target))
	nums = []int{3, 3, 3, 4, 3, 3, 3, 3, 3, 3}
	target = 4
	fmt.Println(search(nums, target))
	nums = []int{3}
	target = 3
	fmt.Println(search(nums, target))
}
