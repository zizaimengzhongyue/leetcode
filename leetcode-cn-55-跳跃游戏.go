package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	max := nums[0] - 1
	for i := 0; i < len(nums); i++ {
		if (i + nums[i]) > max {
			max = i + nums[i]
		}
		if max <= i && i != len(nums)-1 {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))
	nums = []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}
