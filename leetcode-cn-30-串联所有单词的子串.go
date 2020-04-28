package main

import (
	"fmt"
)

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	ln := len(words[0])

	ans := []int{}
	for i := 0; i < len(s); i++ {
		if i+ln*len(words) > len(s) {
			break
		}
		bag := map[string]int{}
		for _, v := range words {
			bag[v]++
		}
		flag := true
		for j := i; j < i+ln*len(words); j += ln {
			v, ok := bag[s[j:j+ln]]
			if !ok {
				flag = false
				break
			}
			if v == 0 {
				flag = false
				break
			}
			bag[s[j:j+ln]]--
		}
		if flag {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	fmt.Println(findSubstring(s, words))
	s = "wordgoodgoodgoodbestword"
	words = []string{"word", "good", "best", "good"}
	fmt.Println(findSubstring(s, words))
	s = "wor"
	words = []string{"word", "good", "word", "word"}
	fmt.Println(findSubstring(s, words))
	s = ""
	words = []string{"word", "good", "word", "word"}
	fmt.Println(findSubstring(s, words))
	s = "wordgoodgoodgoodbestword"
	words = []string{"word", "good", "best", "good"}
	fmt.Println(findSubstring(s, words))
}

/*"wordgoodgoodgoodbestword"
["word","good","best","good"]*/
