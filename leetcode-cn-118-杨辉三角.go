package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	ans := [][]int{[]int{1}}
	for i := 1; i < numRows; i++ {
		tmp := []int{1}
		for j := 1; j <= i; j++ {
			if j < len(ans[i-1]) {
				tmp = append(tmp, ans[i-1][j-1]+ans[i-1][j])
			} else {
				tmp = append(tmp, ans[i-1][j-1])
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}

func main() {
	num := 5
	fmt.Println(generate(num))
}
