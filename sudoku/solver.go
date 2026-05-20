package sudoku

import (
	"fmt"
	"time"
)

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

func isNumUsed(num byte, i int, j int, row [9][10]bool, col [9][10]bool, box [9][10]bool) bool {
	digit := int(num - '0')
	return !row[i][digit] && !col[j][digit] && !box[(j/3)+(i/3)*3][digit]
}

func solveSudoku(board [][]byte) {
	checker := copyBoard(board)

	// these arrays are used to help us determine if we have entered a given digit into
	// the board. row[i][digit] for example tells us if `digit` has been used in row i
	row := [9][10]bool{}
	col := [9][10]bool{}
	box := [9][10]bool{}

	// initialize the row, col and box arrays with the values that are already in them
	for i := range 9 {
		for j := range 9 {
			if board[i][j] != '.' {
				digit := int(board[i][j] - '0')
				row[i][digit] = true
				col[j][digit] = true
				box[(j/3)+(i/3)*3][digit] = true
			}
		}
	}

	v := 0
	for v < 81 {
		i, j := getIJ(v)
		if checker[i][j] == '.' {
			switch board[i][j] {
			case '.':
				board[i][j] = '1'
			case '9':
				digit := int(board[i][j] - '0')
				row[i][digit] = false
				col[j][digit] = false
				box[(j/3)+(i/3)*3][digit] = false
				board[i][j] = '.'
				v -= 1
				i, j = getIJ(v)
				for checker[i][j] != '.' {
					v -= 1
					i, j = getIJ(v)
				}
				continue
			default:
				digit := int(board[i][j] - '0')
				row[i][digit] = false
				col[j][digit] = false
				box[(j/3)+(i/3)*3][digit] = false
				board[i][j] += 1
			}

			isValid := false
			for board[i][j] <= '9' {
				isValid = isNumUsed(board[i][j], i, j, row, col, box)
				if isValid {
					digit := int(board[i][j] - '0')
					row[i][digit] = true
					col[j][digit] = true
					box[(j/3)+(i/3)*3][digit] = true
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

	s := time.Now()
	board = [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '9', '.', '.', '1', '.', '.', '3', '.'},
		{'.', '.', '6', '.', '2', '.', '7', '.', '.'},
		{'.', '.', '.', '3', '.', '4', '.', '.', '.'},
		{'2', '1', '.', '.', '.', '.', '.', '9', '8'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '2', '5', '.', '6', '4', '.', '.'},
		{'.', '8', '.', '.', '.', '.', '.', '1', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
	solveSudoku(board)
	e := time.Now()
	elapsed := e.Sub(s)
	fmt.Printf("%v elapsed\n", elapsed)
	printBoard(board)

}
