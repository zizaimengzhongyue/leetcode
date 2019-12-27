package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	dict := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		b := []byte(strs[i])
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		dict[string(b)] = append(dict[string(b)], strs[i])
	}
	ans := [][]string{}
	for _, v := range dict {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
