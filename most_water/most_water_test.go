package mostwater

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMostWater(t *testing.T) {

	res := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	require.Equal(t, 49, res)

}
