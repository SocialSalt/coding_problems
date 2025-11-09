package dynamicprogramming

func uniquePaths(m int, n int) int {

	grid := make([][]int, m)
	for i := range m {
		grid[i] = make([]int, n)
	}

	for i := range m {
		for j := range n {
			if i == 0 && j == 0 {
				grid[i][j] = 1
			}
			if i < m-1 {
				grid[i+1][j] += grid[i][j]
			}
			if j < n-1 {
				grid[i][j+1] += grid[i][j]
			}
		}
	}
	return grid[m-1][n-1]

}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	grid := make([][]int, m)
	for i := range m {
		grid[i] = make([]int, n)
	}

	for i := range m {
		for j := range n {
			if i == 0 && j == 0 && obstacleGrid[i][j] == 0 {
				grid[i][j] = 1
			}
			if i < m-1 && obstacleGrid[i+1][j] == 0 {
				grid[i+1][j] += grid[i][j]
			}
			if j < n-1 && obstacleGrid[i][j+1] == 0 {
				grid[i][j+1] += grid[i][j]
			}
		}
	}
	return grid[m-1][n-1]

}
