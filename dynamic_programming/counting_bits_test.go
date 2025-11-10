package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountBits(t *testing.T) {
	var res []int
	res = countBits(2)
	require.Equal(t, []int{0, 1, 1}, res)

	res = countBits(5)
	require.Equal(t, []int{0, 1, 1, 2, 1, 2}, res)
}
