package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntegerBreak(t *testing.T) {
	var res int
	res = integerBreak(2)
	require.Equal(t, 1, res)

	res = integerBreak(5)
	require.Equal(t, 6, res)

	res = integerBreak(3)
	require.Equal(t, 2, res)

	res = integerBreak(8)
	require.Equal(t, 18, res)

	res = integerBreak(10)
	require.Equal(t, 36, res)
}
