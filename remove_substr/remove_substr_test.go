package removesubstr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaximumGain(t *testing.T) {
	s := "cdbcbbaaabab"
	x := 4
	y := 5
	result := maximumGain(s, x, y)
	require.Equal(t, 19, result)
}
