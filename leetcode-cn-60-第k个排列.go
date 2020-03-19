package main

import "fmt"

func getNstValue(n int, mark []int) uint {
	index := 0
	for k, v := range mark {
		if v == 0 {
			index = index + 1
		}
		if index == n {
			mark[k] = 1
			return uint(k + 1)
		}
	}
	return uint(index)
}

func dfs(n int, k int, dp []int, mark []int) []byte {
	ans := []byte{}
	tmp := 0
	if n-1 != 0 {
		tmp = (k - 1) / dp[n-1]
	}
	v := getNstValue(tmp+1, mark)
	ans = append(ans, byte('0' + v))
	k = k - tmp*dp[n-1]
	if n-1 != 0 {
		ans = append(ans, dfs(n-1, k, dp, mark)...)
	}
	return ans
}

func getPermutation(n int, k int) string {
	dp := []int{}
	mark := make([]int, n, n)
	tmp := 1
	dp = append(dp, 0)
	for i := 1; i <= n; i++ {
		tmp = tmp * i
		dp = append(dp, tmp)
	}
	ans := dfs(n, k, dp, mark)
    return string(ans)
}

func main() {
	n := 3
	k := 3
	fmt.Println(getPermutation(n, k))
	n = 3
	k = 6
	fmt.Println(getPermutation(n, k))
	n = 1
	k = 1
	fmt.Println(getPermutation(n, k))
	n = 5
	k = 120
	fmt.Println(getPermutation(n, k))
}
