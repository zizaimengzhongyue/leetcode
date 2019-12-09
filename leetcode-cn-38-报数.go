package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	ans := []byte{'1'}
	for i := n - 1; i > 0; i-- {
		temp := []byte{}
		ln := len(ans)
		t := ans[0]
		count := 1
		for j := 1; j < ln; j++ {
			if ans[j] == t {
				count++
				continue
			} else {
                tt := strconv.Itoa(count)
				temp = append(temp, tt...)
				temp = append(temp, t)
				t = ans[j]
				count = 1
			}
		}
        tt := strconv.Itoa(count)
		temp = append(temp, tt...)
		temp = append(temp, t)
		ans = temp
	}
	return string(ans)
}

func main() {
	fmt.Println(countAndSay(1))
    fmt.Println(countAndSay(2))
    fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
	fmt.Println(countAndSay(10))
}
