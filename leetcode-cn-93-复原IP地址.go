package main

import (
	"fmt"
	"strconv"
)

func dfs(s string, begin int, tmp string, ln int, ans []string, step int) []string {
	if step == 4 && begin == len(s) {
		ans = append(ans, tmp[0:len(tmp)-1])
		return ans
	}
	if step > 4 || begin > len(s)-1 {
		return ans
	}
	tmp += (s[begin:begin+1] + ".")
	ans = dfs(s, begin+1, tmp, ln+2, ans, step+1)
	tmp = tmp[:ln+1]
	if s[begin] == '0' {
		goto Return
	}
	if begin+1 < len(s) {
		tmp += (s[begin:begin+2] + ".")
		ans = dfs(s, begin+2, tmp, ln+3, ans, step+1)
		tmp = tmp[:ln+1]
	}
	if begin+2 < len(s) {
		t, _ := strconv.Atoi(s[begin : begin+3])
		if t > 255 {
			goto Return
		}
		tmp += (s[begin:begin+3] + ".")
		ans = dfs(s, begin+3, tmp, ln+4, ans, step+1)
		tmp = tmp[:ln+1]
	}
Return:
	return ans
}

func restoreIpAddresses(s string) []string {
	ans := []string{}
	if len(s) < 4 || len(ans) > 12 {
		return ans
	}
	return dfs(s, 0, "", -1, ans, 0)
}

func main() {
	str := "25525511135"
	fmt.Println(restoreIpAddresses(str))
	str = "0000"
	fmt.Println(restoreIpAddresses(str))
	str = "25525511135"
	fmt.Println(restoreIpAddresses(str))
	str = "25525511135"
	fmt.Println(restoreIpAddresses(str))
}
