package main

import "fmt"

func dfs(n int, k int, begin int, tmp []int, index int, ans [][]int) [][]int {
	if index == k {
		return append(ans, append([]int{}, tmp...))
	}
	for i := begin; i <= n; i++ {
		tmp[index] = i
		ans = dfs(n, k, i+1, tmp, index+1, ans)
	}
	return ans
}

func combine(n int, k int) [][]int {
	ans := [][]int{}
	if n == 0 || k == 0 || n < k {
		return ans
	}
	for i := 1; i <= n; i++ {
		tmp := make([]int, k, k)
		tmp[0] = i
		ans = dfs(n, k, i+1, tmp, 1, ans)
	}
	return ans
}

func main() {
	n := 4
	k := 2
	fmt.Println(combine(n, k))
	n = 2
	k = 2
	fmt.Println(combine(n, k))
	n = 2
	k = 0
	fmt.Println(combine(n, k))
}
