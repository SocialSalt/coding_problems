package greedyproblems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanJump(t *testing.T) {
	res := canJump([]int{2, 3, 1, 1, 4})
	require.True(t, res)

	res = canJump([]int{3, 2, 1, 0, 4})
	require.False(t, res)
}

func TestJump(t *testing.T) {
	var res int
	res = jump([]int{2, 3, 1, 1, 4})
	require.Equal(t, 2, res)

	res = jump([]int{2, 3, 0, 1, 4})
	require.Equal(t, 2, res)

	res = jump([]int{2, 1})
	require.Equal(t, 1, res)

	res = jump([]int{9, 8, 2, 2, 0, 2, 2, 0, 4, 1, 5, 7, 9, 6, 6, 0, 6, 5, 0, 5})
	require.Equal(t, 3, res)

	res = jump([]int{0})
	require.Equal(t, 0, res)
}
