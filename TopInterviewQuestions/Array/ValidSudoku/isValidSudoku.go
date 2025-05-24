package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	// column check
	for i := 0; i < 9; i++ {
		validator := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] == 0x2e {
				continue
			}
			_, ok := validator[board[i][j]]
			if ok {
				return false
			} else {
				validator[board[i][j]] = true
			}
		}
	}

	// row check
	for j := 0; j < 9; j++ {
		validator := make(map[byte]bool)
		for i := 0; i < 9; i++ {
			if board[i][j] == 0x2e {
				continue
			}
			_, ok := validator[board[i][j]]
			if ok {
				return false
			} else {
				validator[board[i][j]] = true
			}
		}
	}

	// 3x3
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			validator := make(map[byte]bool)
			for n := 0; n < 3; n++ {
				for m := 0; m < 3; m++ {
					if board[i+n][j+m] == 0x2e {
						continue
					}
					_, ok := validator[board[i+n][j+m]]
					if ok {
						return false
					} else {
						validator[board[i+n][j+m]] = true
					}
				}
			}
		}
	}

	return true
}

func main() {
	board := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '3', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'1', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	res := isValidSudoku(board)
	fmt.Print(res)
}
