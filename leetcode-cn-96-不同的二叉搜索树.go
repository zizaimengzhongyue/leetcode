package main

import (
	"fmt"
)

func numTrees(n int) int {
	ans := 1
	for i := 2; i <= n; i++ {
		ans = ans*2 + ans - 1
	}
	return ans
}

func main() {
	num := 3
	fmt.Println(numTrees(num))
	num = 2
	fmt.Println(numTrees(num))
	num = 5
	fmt.Println(numTrees(num))
}
