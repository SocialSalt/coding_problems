package greedyproblems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinMoves(t *testing.T) {
	nums := []int{28, 50, 76, 80, 64, 30, 32, 84, 53, 8}
	limit := 84

	moves := minMoves(nums, limit)
	require.Equal(t, 4, moves)
}
