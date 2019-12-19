package main

import "fmt"

func check(board [][]byte, x int, y int) bool {
	keys := map[byte]bool{'1': false, '2': false, '3': false, '4': false, '5': false, '6': false, '7': false, '8': false, '9': false}
	for j := 0; j < 9; j++ {
		if board[x][j] != '.' {
            if keys[board[x][j]] {
                return false
            }
			keys[board[x][j]] = true
		}
	}
	keys = map[byte]bool{'1': false, '2': false, '3': false, '4': false, '5': false, '6': false, '7': false, '8': false, '9': false}
	for i := 0; i < 9; i++ {
		if board[i][y] != '.' {
            if keys[board[i][y]] {
                return false
            }
			keys[board[i][y]] = true
		}
	}
	if x%3 != 0 {
		x--
	}
	if x%3 != 0 {
		x--
	}
	if y%3 != 0 {
		y--
	}
	if y%3 != 0 {
		y--
	}
	keys = map[byte]bool{'1': false, '2': false, '3': false, '4': false, '5': false, '6': false, '7': false, '8': false, '9': false}
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			if board[i][j] != '.' {
                if keys[board[i][j]] {
                    return false
                }
				keys[board[i][j]] = true
			}
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				if !check(board, i, j) {
                    return false
                }
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
	board = [][]byte{
		[]byte{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))

	board = [][]byte{
		[]byte{'.', '8', '7', '6', '5', '4', '3', '2', '1'},
		[]byte{'2', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'3', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'4', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'5', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'8', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'9', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Println(isValidSudoku(board))
}
