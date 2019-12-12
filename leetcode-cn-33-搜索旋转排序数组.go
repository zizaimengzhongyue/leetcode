package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if r-l <= 1 {
			if nums[l] == target {
				return l
			}
			if nums[r] == target {
				return r
			}
			return -1
		}
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target && target <= nums[r] {
			l = mid
            continue
		}
		if nums[mid] < target && target > nums[r] && nums[mid] > nums[l] {
			l = mid
            continue
		}
        if nums[mid] < target && target > nums[r] && nums[mid] < nums[l] {
            r = mid
            continue
        }
		if nums[mid] > target && target >= nums[l] {
			r = mid
            continue
		}
		if nums[mid] > target && target < nums[l] && nums[mid] < nums[l] {
			r = mid
            continue
		}
        if nums[mid] > target && target < nums[l] && nums[mid] > nums[l] {
            l = mid
        }
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))
	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 3
	fmt.Println(search(nums, target))
	nums = []int{3, 5, 1}
	target = 3
	fmt.Println(search(nums, target))
	nums = []int{5, 1, 3}
	target = 3
	fmt.Println(search(nums, target))
	nums = []int{5, 1, 3}
	target = 5
	fmt.Println(search(nums, target))
	nums = []int{4, 5, 6, 7, 8, 1, 2, 3}
	target = 8
	fmt.Println(search(nums, target))
}
