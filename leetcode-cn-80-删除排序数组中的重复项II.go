package main

import (
	"fmt"
)

func dfs(nums []int, index int, cur int, cnt int, shift int) int {
    if index + shift < len(nums) {
        nums[index] = nums[index + shift]
    } else {
        return shift
    }
    if nums[index] == cur {
        cnt++
    } else {
        cur = nums[index]
        cnt = 1
    }
    if cnt > 2 {
        shift++
        cnt = 2
    } else {
        index++
    }
    return dfs(nums, index, cur, cnt, shift)
}

func removeDuplicates(nums []int) int {
	return len(nums) - dfs(nums, 0, -1, 0, 0)
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
	nums = []int{1, 1, 1}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
	nums = []int{1}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
	nums = []int{}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
	nums = []int{1, 1, 1, 1}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
	nums = []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
