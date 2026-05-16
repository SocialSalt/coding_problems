package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxBalancedShipments(t *testing.T) {

	weight := []int{2, 5, 1, 4, 3}
	num := maxBalancedShipments(weight)
	require.Equal(t, 2, num)

	weight = []int{4, 4}
	num = maxBalancedShipments(weight)
	require.Equal(t, 0, num)
}
