package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	ln := len(s)
	ans := 0
	for i := 0; i < ln; i++ {
		left := 0
		right := 0
		for j := i; j < ln; j++ {
			if s[j] == '(' {
				left++
			} else {
				right++
			}
			if left == right {
				if ans < right*2 {
					ans = right * 2
				}
			}
			if left < right {
				left = 0
				right = 0
			}
		}
	}
	return ans
}

func main() {
	str := ")()())"
	fmt.Println(longestValidParentheses(str))
	str = "()(()"
	fmt.Println(longestValidParentheses(str))
	str = "(()(((()"
	fmt.Println(longestValidParentheses(str))
	str = "(()"
	fmt.Println(longestValidParentheses(str))
}
