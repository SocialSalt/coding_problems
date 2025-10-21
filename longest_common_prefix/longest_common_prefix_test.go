package longestcommonprefix

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestCommonPrefix(t *testing.T) {
	res := longestCommonPrefix([]string{"flower", "flow", "flight"})
	require.Equal(t, "fl", res)
}
