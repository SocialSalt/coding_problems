package sudoku

func checkValidEntry(board [][]byte, i int, j int) bool {
	if valid := sectionIsValid(board[i]); !valid {
		return false
	}
	col := make([]byte, 9)
	for k := range 9 {
		col[k] = board[k][j]
	}
	if valid := sectionIsValid(col); !valid {
		return false
	}
	r, c := (i/3)*3, (j/3)*3
	for k := range 3 {
		for l := range 3 {
			col[k*3+l] = board[r+k][c+k]
		}
	}
	if valid := sectionIsValid(col); !valid {
		return false
	}
	return true
}

func getPossibleEntries(board [][]byte, i int, j int) []byte {
	return nil
}

type Cell struct {
	I               int
	J               int
	PossibleEntries []byte
}

func solveSudoku(board [][]byte) {
	cells := []Cell{}
	for i := range 9 {
		for j := range 9 {
			if board[i][j] == '.' {
				possibleEntries[i][j] = []byte{}
			}
		}
	}

}
