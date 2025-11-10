package dynamicprogramming

func longestIncreasingPath(matrix [][]int) int {
	height := len(matrix)
	width := len(matrix[0])
	longestPath := 0

	longestPaths := make([][]int, height)
	for i := range height {
		longestPaths[i] = make([]int, width)
	}

	var findLongest func(x int, y int) int
	findLongest = func(x int, y int) int {
		if longestPaths[x][y] > 0 {
			return longestPaths[x][y]
		}
		l := 1
		for _, dir := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			i_ := x + dir[0]
			j_ := y + dir[1]
			if i_ > height-1 || j_ > width-1 || i_ < 0 || j_ < 0 {
				continue
			}
			if matrix[i_][j_] > matrix[x][y] {
				l = max(l, 1+findLongest(i_, j_))
			}
		}
		longestPaths[x][y] = l
		return l
	}

	for i := range matrix {
		for j := range matrix[i] {
			longestPath = max(longestPath, findLongest(i, j))
		}
	}
	return longestPath
}
