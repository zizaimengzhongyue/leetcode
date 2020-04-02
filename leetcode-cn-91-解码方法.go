package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	ln := len(s)
	if ln == 0 {
		return 0
	} else if s[0] == '0' {
		return 0
	}
	dp := make([]int, ln, ln)
	dp[0] = 1
	for i := 1; i < ln; i++ {
		t1, _ := strconv.Atoi(s[i : i+1])
		if t1 > 0 {
			dp[i] += dp[i-1]
		}
		if s[i-1] == '0' {
			continue
		}
		if i-2 >= 0 {
			t2, _ := strconv.Atoi(s[i-1 : i+1])
			if t2 > 0 && t2 <= 26 {
				dp[i] += dp[i-2]
			}
		} else if i-2 < 0 {
			t2, _ := strconv.Atoi(s[i-1 : i+1])
			if t2 > 0 && t2 <= 26 {
				dp[i] += 1
			}
		}
	}
	return dp[ln-1]
}

func main() {
	str := "12"
	fmt.Println(numDecodings(str))
	str = "226"
	fmt.Println(numDecodings(str))
	str = "10"
	fmt.Println(numDecodings(str))
	str = "100"
	fmt.Println(numDecodings(str))
	str = "101"
	fmt.Println(numDecodings(str))
}
