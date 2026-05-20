package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindPrefixCommon(t *testing.T) {
	A := []int{1, 3, 2, 4}
	B := []int{3, 1, 2, 4}
	C := findThePrefixCommonArray(A, B)
	require.Equal(t, []int{0, 2, 3, 4}, C)

}
