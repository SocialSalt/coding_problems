package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCoinChange(t *testing.T) {
	var res int
	res = coinChange([]int{1, 2, 5}, 11)
	require.Equal(t, 3, res)

	res = coinChange([]int{2}, 3)
	require.Equal(t, -1, res)

	res = coinChange([]int{1}, 0)
	require.Equal(t, 0, res)

	res = coinChange([]int{1, 214748}, 2)
	require.Equal(t, 2, res)

	res = coinChange([]int{2, 5, 10, 1}, 27)
	require.Equal(t, 4, res)

	res = coinChange([]int{186, 419, 83, 408}, 6249)
	require.Equal(t, 20, res)
}
