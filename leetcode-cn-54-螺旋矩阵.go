package main

import (
	"fmt"
)

type step struct {
	target int
	count  int
}

func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return ans
	}
	m := len(matrix)
	n := len(matrix[0])
	if m == 1 && n == 1 {
		ans = append(ans, matrix[0][0])
		return ans
	}
	i := 0
	j := 0
	s := step{target: 0, count: 0}
	for m != 0 && n != 0 {
		ans = append(ans, matrix[i][j])
		s.count++
		if s.target%2 == 0 && s.count >= n {
			s.target = (s.target + 1) % 4
			s.count = 0
			m--
		} else if s.target%2 == 1 && s.count >= m {
			s.target = (s.target + 1) % 4
			s.count = 0
			n--
		}
		if s.target == 0 {
			j++
		}
		if s.target == 1 {
			i++
		}
		if s.target == 2 {
			j--
		}
		if s.target == 3 {
			i--
		}
	}
	return ans
}

func main() {
	//nums:= [][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}
	//nums := [][]int{{1, 2}, {3, 4}}
	nums := [][]int{{1}}
	fmt.Println(spiralOrder(nums))
}
