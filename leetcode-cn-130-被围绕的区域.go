package main

import (
	"fmt"
)

type Step struct {
	I int
	J int
}

func initMark(m int, n int) [][]bool {
	mark := make([][]bool, m)
	for i := 0; i < m; i++ {
		mark[i] = make([]bool, n)
	}
	return mark
}

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m := len(board)
	n := len(board[0])
	mark := initMark(m, n)
	q := []Step{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && (i == 0 || i == m-1 || j == 0 || j == n-1) {
				q = append(q, Step{I: i, J: j})
				mark[i][j] = true
			}
		}
	}

	head := 0
	tail := len(q)
	for head < tail {
		cur := q[head]
		head++
		if cur.I+1 < m && cur.J < n && board[cur.I+1][cur.J] == 'O' && !mark[cur.I+1][cur.J] {
			mark[cur.I+1][cur.J] = true
			q = append(q, Step{I: cur.I + 1, J: cur.J})
			tail++
		}
		if cur.I-1 >= 0 && cur.J < n && board[cur.I-1][cur.J] == 'O' && !mark[cur.I-1][cur.J] {
			mark[cur.I-1][cur.J] = true
			q = append(q, Step{I: cur.I - 1, J: cur.J})
			tail++
		}
		if cur.I < m && cur.J+1 < n && board[cur.I][cur.J+1] == 'O' && !mark[cur.I][cur.J+1] {
			mark[cur.I][cur.J+1] = true
			q = append(q, Step{I: cur.I, J: cur.J + 1})
			tail++
		}
		if cur.I < m && cur.J-1 >= 0 && board[cur.I][cur.J-1] == 'O' && !mark[cur.I][cur.J-1] {
			mark[cur.I][cur.J-1] = true
			q = append(q, Step{I: cur.I, J: cur.J - 1})
			tail++
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !mark[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

func main() {
	b := [][]byte{}
	b = append(b, []byte{'X', 'X', 'X', 'X'})
	b = append(b, []byte{'X', 'O', 'O', 'X'})
	b = append(b, []byte{'X', 'X', 'O', 'X'})
	b = append(b, []byte{'X', 'O', 'X', 'X'})
	solve(b)
	fmt.Println(b)

	b = [][]byte{[]byte{'X'}}
	solve(b)
	fmt.Println(b)

	b = [][]byte{}
	b = append(b, []byte{'O', 'O'})
	b = append(b, []byte{'O', 'O'})
	solve(b)
	fmt.Println(b)
}
