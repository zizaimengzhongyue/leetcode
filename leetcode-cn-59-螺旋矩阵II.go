package main

import "fmt"

func generateMatrix(n int) [][]int {
	ans := [][]int{}
	for i := 0; i < n; i++ {
		ans = append(ans, []int{})
		for j := 0; j < n; j++ {
			ans[i] = append(ans[i], 0)
		}
	}
	i := 0
	j := 0
	target := 0
	limit := n
	step := 0
	start := 0
	end := 3
	for v := 1; v <= n*n; v++ {
		ans[i][j] = v
		target++
		if step == 0 {
			j++
		} else if step == 1 {
			i++
		} else if step == 2 {
			j--
		} else if step == 3 {
			i--
		}
		if target >= limit - 1 {
			target = 0
			step = (step + 1) % 4
			start = start + 1
			if start == end {
				limit--
				start = 0
                if end != 2 {
				    end--
                }
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(generateMatrix(1))
	fmt.Println(generateMatrix(2))
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(4))
	fmt.Println(generateMatrix(5))
}
