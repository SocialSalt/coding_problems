package arraysum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test3Sum(t *testing.T) {
	res := threeSum([]int{0, 1, 1})
	require.Equal(t, [][]int{}, res)

	res = threeSum([]int{0, 0, 0})
	require.Equal(t, [][]int{{0, 0, 0}}, res)

	res = threeSum([]int{-1, 0, 1, 2, -1, -4})
	require.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, res)
}

func Test3SumTarget(t *testing.T) {
	res := threeSumClosest([]int{-1, 2, 1, -4}, 1)
	require.Equal(t, 2, res)

	res = threeSumClosest([]int{0, 0, 0}, 1)
	require.Equal(t, 0, res)
}

func Test4Sum(t *testing.T) {
	res := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	require.Equal(t, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}, res)

	res = fourSum([]int{-3, -1, 0, 2, 4, 5}, 2)
	require.Equal(t, [][]int{{-3, -1, 2, 4}}, res)
}
