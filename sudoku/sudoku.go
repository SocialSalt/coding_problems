package sudoku

func sectionIsValid(section []byte) bool {
	found := make(map[byte]struct{}, 9)
	for _, item := range section {
		if item == '.' {
			continue
		}
		_, ok := found[item]
		if ok {
			return false
		}
		found[item] = struct{}{}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	section := make([]byte, 9)
	for i := range 9 {
		is_valid := sectionIsValid(board[i])
		if !is_valid {
			return false
		}
		for j := range 9 {
			section[j] = board[j][i]
		}
		is_valid = sectionIsValid(section)
		if !is_valid {
			return false
		}

		c := (i * 3) % 9
		r := (i / 3) * 3
		for j := range 3 {
			for k := range 3 {
				section[3*j+k] = board[r+j][c+k]
			}
		}
		is_valid = sectionIsValid(section)
		if !is_valid {
			return false
		}
	}

	return true
}
