package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinimumCost(t *testing.T) {
	cost := minimumCost(3, 2, []int{1, 3}, []int{5})
	require.Equal(t, 13, cost)
}
