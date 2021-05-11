package main

import "fmt"

func merge(a1 [][]string, a2 [][]string) [][]string {
	if len(a2) == 0 {
		return a1
	}
	ans := [][]string{}
	for _, v1 := range a1 {
		tmp := []string{}
		for _, v2 := range a2 {
			tmp = append(append([]string{}, v1...), v2...)
			ans = append(ans, tmp)
		}
	}
	return ans
}

func isValid(s string) bool {
	ln := len(s)
	for i := 0; i < ln/2; i++ {
		if s[i] != s[ln-1-i] {
			return false
		}
	}
	return true
}

func partition(s string) [][]string {
	if len(s) == 1 {
		return [][]string{[]string{s}}
	}
	ans := [][]string{}
	for i := 1; i <= len(s); i++ {
		if !isValid(s[0:i]) {
			continue
		}
		tmp := [][]string{}
		if i < len(s) {
			tmp = partition(s[i:])
		}
		ans = append(ans, merge([][]string{[]string{s[0:i]}}, tmp)...)
	}
	return ans
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
	s = "a"
	fmt.Println(partition(s))
	s = "abcdcba"
	fmt.Println(partition(s))
	s = ""
	fmt.Println(partition(s))
	s = "bb"
	fmt.Println(partition(s))
}
