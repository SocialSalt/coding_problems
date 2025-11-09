package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniquePaths(t *testing.T) {
	var res int
	res = uniquePaths(3, 7)
	require.Equal(t, 28, res)
}

func TestUniquePathsWithObsticle(t *testing.T) {
	var res int
	res = uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})
	require.Equal(t, 2, res)

}
