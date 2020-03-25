package main

import "fmt"

func bfs(board [][]byte, i int, j int, bts []byte, k int, mark [][]int) bool {
	if board[i][j] != bts[k] {
		return false
	}
	if k == len(bts)-1 {
		return true
	}
	mark[i][j] = 1
	m := len(board)
	n := len(board[i])
	if i-1 >= 0 && mark[i-1][j] != 1 {
		if bfs(board, i-1, j, bts, k+1, mark) {
			return true
		}
	}
	if i+1 < m && mark[i+1][j] != 1 {
		if bfs(board, i+1, j, bts, k+1, mark) {
			return true
		}
	}
	if j-1 >= 0 && mark[i][j-1] != 1 {
		if bfs(board, i, j-1, bts, k+1, mark) {
			return true
		}
	}
	if j+1 < n && mark[i][j+1] != 1 {
		if bfs(board, i, j+1, bts, k+1, mark) {
			return true
		}
	}
	mark[i][j] = 0
	return false
}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	mark := make([][]int, len(board), len(board))
	for i := 0; i < len(board); i++ {
		mark[i] = make([]int, len(board[i]), len(board[i]))
	}
	bts := []byte(word)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if bfs(board, i, j, bts, 0, mark) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	fmt.Println(exist(board, word))
	word = "SEE"
	fmt.Println(exist(board, word))
	word = "ABCB"
	fmt.Println(exist(board, word))
	word = ""
	fmt.Println(exist(board, word))
	board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}
	word = "ABCESEEEFS"
	fmt.Println(exist(board, word))
}
