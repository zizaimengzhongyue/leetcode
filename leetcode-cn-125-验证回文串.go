package main

import (
	"fmt"
	"strings"
)

func isByte(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= '0' && b <= '9'
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	ns := []byte{}
	for i := 0; i < len(s); i++ {
		if !isByte(s[i]) {
			continue
		}
		ns = append(ns, s[i])
	}
	i := 0
	j := len(ns) - 1
	for i < j {
		if ns[i] != ns[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	a := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(a))
	a = "OP"
	fmt.Println(isPalindrome(a))
}
