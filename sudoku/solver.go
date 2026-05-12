package sudoku

import (
	"fmt"
)

func checkValidEntry(board [][]byte, i int, j int) bool {
	if board[i][j] == '.' {
		return false
	}
	// check row
	if valid := sectionIsValid(board[i]); !valid {
		return false
	}

	// check col
	col := make([]byte, 9)
	for k := range 9 {
		col[k] = board[k][j]
	}
	if valid := sectionIsValid(col); !valid {
		return false
	}

	// check square
	r, c := (i/3)*3, (j/3)*3
	for k := range 3 {
		for l := range 3 {
			col[k*3+l] = board[r+k][c+l]
		}
	}
	if valid := sectionIsValid(col); !valid {
		return false
	}
	return true
}

func getIJ(v int) (int, int) {
	j := v % 9
	i := v / 9
	return i, j
}

func printBoard(board [][]byte) {
	for i := range 9 {
		for j := range 9 {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println()
}

func copyBoard(board [][]byte) [][]byte {
	suggestion := make([][]byte, 9)
	for i := range 9 {
		suggestion[i] = make([]byte, 9)
	}
	for i := range 9 {
		for j := range 9 {
			suggestion[i][j] = board[i][j]
		}
	}
	return suggestion

}

func solveSudoku(board [][]byte) {
	checker := copyBoard(board)
	v := 0
	for v < 81 {
		i, j := getIJ(v)
		if checker[i][j] == '.' {
			switch board[i][j] {
			case '.':
				board[i][j] = '1'
			case '9':
				board[i][j] = '.'
				v -= 1
				i, j = getIJ(v)
				for checker[i][j] != '.' {
					v -= 1
					i, j = getIJ(v)
				}
				continue
			default:
				board[i][j] += 1
			}

			isValid := false
			for board[i][j] <= '9' {
				isValid = checkValidEntry(board, i, j)
				if isValid {
					break
				}
				board[i][j] += 1
			}
			//backtrack
			if !isValid {
				board[i][j] = '.'
				v -= 1
				i, j = getIJ(v)
				for checker[i][j] != '.' {
					v -= 1
					i, j = getIJ(v)
				}
			} else {
				v += 1
				i, j = getIJ(v)
			}
		} else {
			v += 1
			i, j = getIJ(v)
		}
	}
}

func RunSudoku() {

	var board [][]byte

	board = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	printBoard(board)
}
