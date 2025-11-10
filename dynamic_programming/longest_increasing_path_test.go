package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestIncreasingPath(t *testing.T) {
	var res int
	res = longestIncreasingPath([][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	})
	require.Equal(t, 4, res)

	res = longestIncreasingPath([][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	})
	require.Equal(t, 4, res)

	res = longestIncreasingPath([][]int{
		{7, 6, 1, 1},
		{2, 7, 6, 0},
		{1, 3, 5, 1},
		{6, 6, 3, 2},
	})
	require.Equal(t, 7, res)

	res = longestIncreasingPath([][]int{
		{1, 2},
	})
	require.Equal(t, 2, res)
}
