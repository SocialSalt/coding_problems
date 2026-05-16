package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinimumEffort(t *testing.T) {
	tasks := [][]int{{2, 4}, {1, 2}, {4, 8}}
	energy := minimumEffort(tasks)
	require.Equal(t, 8, energy)
}
